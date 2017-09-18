package replpkg

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"

	// Importing this package installs Import as go/types.DefaultImport.
	"golang.org/x/tools/imports"

	"github.com/motemen/go-quickfix"
)

const printerName = "__gophernotes"

// Session encodes info about the current REPL session.
type Session struct {
	FilePath       string
	File           *ast.File
	Fset           *token.FileSet
	Types          *types.Config
	TypeInfo       types.Info
	ExtraFilePaths []string
	ExtraFiles     []*ast.File

	mainBody         *ast.BlockStmt
	storedBodyLength int
}

const oldStdout = "__oldStdOut"
const initialSourceTemplate = `
package main

import (
	%q
	"os"
)

func ` + printerName + `(xx ...interface{}) {
	for _, x := range xx {
		%s
	}
}

func main() {
    // Discard any print statements. This should be restore before the current cell is executed
	` + oldStdout + ` := os.Stdout
	os.Stdout = nil
}
`

// printerPkgs is a list of packages that provides
// pretty printing function.
var printerPkgs = []struct {
	path string
	code string
}{
	{"github.com/k0kubun/pp", `pp.Println(x)`},
	{"github.com/davecgh/go-spew/spew", `spew.Printf("%#v\n", x)`},
	{"fmt", `fmt.Printf("%#v\n", x)`},
}

// NewSession initiates a new REPL
func NewSession() (*Session, error) {

	s := &Session{
		Fset: token.NewFileSet(),
		Types: &types.Config{
			Importer: importer.Default(),
		},
	}

	var err error
	s.FilePath, err = tempFile()
	if err != nil {
		return nil, err
	}

	var initialSource string
	for _, pp := range printerPkgs {
		_, err := importer.Default().Import(pp.path)
		if err == nil {
			initialSource = fmt.Sprintf(initialSourceTemplate, pp.path, pp.code)
			break
		}
		debugf("could not import %q: %s", pp.path, err)
	}
	if initialSource == "" {
		return nil, fmt.Errorf("Could not load pretty printing package")
	}

	s.File, err = parser.ParseFile(s.Fset, "gophernotes_session.go", initialSource, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	s.mainBody = s.mainFunc().Body

	return s, nil
}

func (s *Session) mainFunc() *ast.FuncDecl {
	return s.File.Scope.Lookup("main").Decl.(*ast.FuncDecl)
}

// Run calls "go run" with appropriate files appended.
func (s *Session) Run() ([]byte, bytes.Buffer, error) {
	f, err := os.Create(s.FilePath)
	if err != nil {
		return nil, bytes.Buffer{}, err
	}

	err = printer.Fprint(f, s.Fset, s.File)
	if err != nil {
		return nil, bytes.Buffer{}, err
	}

	return goRun(append(s.ExtraFilePaths, s.FilePath))
}

// tempFile prepares the temporary session file for the REPL.
func tempFile() (string, error) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "gophernotes_session.go"), nil
}

func goRun(files []string) ([]byte, bytes.Buffer, error) {

	var stderr bytes.Buffer

	args := append([]string{"run"}, files...)
	debugf("go %s", strings.Join(args, " "))
	cmd := exec.Command("go", args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = &stderr
	out, err := cmd.Output()

	return out, stderr, err
}

// If a valid expression, wraps the expression in a print statement to display
// the output as expected in a notebook context
func (s *Session) evalExpr(in string) (ast.Expr, error) {
	expr, err := parser.ParseExpr(in)
	if err != nil {
		return nil, err
	}

	stmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  ast.NewIdent(printerName),
			Args: []ast.Expr{expr},
		},
	}

	s.appendStatements(stmt)

	return expr, nil
}

func isNamedIdent(expr ast.Expr, name string) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == name
}

func (s *Session) evalStmt(in string, noPrint bool) error {
	src := fmt.Sprintf("package P; func F() { %s }", in)
	f, err := parser.ParseFile(s.Fset, "stmt.go", src, parser.Mode(0))

	// If it can't be parsed as a statement then the last option is to include it as a function
	if err != nil {
		return s.evalFunction(in)
	}

	enclosingFunc := f.Scope.Lookup("F").Decl.(*ast.FuncDecl)
	stmts := enclosingFunc.Body.List

	if len(stmts) > 0 {

		debugf("evalStmt :: %s", showNode(s.Fset, stmts))
		lastStmt := stmts[len(stmts)-1]

		// print last assigned/defined values
		if !noPrint {
			if assign, ok := lastStmt.(*ast.AssignStmt); ok {
				vs := []ast.Expr{}
				for _, v := range assign.Lhs {
					if !isNamedIdent(v, "_") {
						vs = append(vs, v)
					}
				}

				if len(vs) > 0 {
					printLastValues := &ast.ExprStmt{
						X: &ast.CallExpr{
							Fun:  ast.NewIdent(printerName),
							Args: vs,
						},
					}
					stmts = append(stmts, printLastValues)
				}
			}
		}
	}

	s.appendStatements(stmts...)

	return nil
}

func (s *Session) evalFunction(in string) error {
	// try to import this as a proxy function and correct for any imports
	appendForImport := `package main


			`

	proxyFilePath := string(filepath.Dir(s.FilePath)) + "/func_proxy.go"

	proxyFile, err := os.Create(proxyFilePath)
	if err != nil {
		return err
	}

	_, err = proxyFile.Write([]byte(appendForImport + in))
	if err != nil {
		return err
	}
	proxyFile.Close()

	b := new(bytes.Buffer)
	cmd := exec.Command("goimports", "-w", proxyFilePath)
	cmd.Stdout = b
	cmd.Stderr = b
	err = cmd.Run()

	if err != nil {
		err = errors.New(err.Error() + "\n" + b.String())
		return err
	}

	funcProxy, err := ioutil.ReadFile(proxyFilePath)
	if err != nil {
		return err
	}

	// Parse the proxy
	parsedProxyFile, err := parser.ParseFile(s.Fset, "func_proxy.go", funcProxy, parser.Mode(0))
	if err != nil {
		return err
	}

	return s.mergeFunctionDeclarations(parsedProxyFile.Decls)
}

func (s *Session) mergeFunctionDeclarations(newFunctionDecls []ast.Decl) error {
	newFuncs := make(map[string]*ast.FuncDecl)
	for _, decl := range newFunctionDecls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcName := funcDecl.Name.Name
			if funcName == "main" {
				return errors.New("Cannot declare main function")
			} else if funcName == printerName {
				return errors.New("Cannot overwrite printer magic function " + printerName)
			}
			newFuncs[funcName] = funcDecl
		}
	}

	currentDecls := s.File.Decls
	for i, decl := range currentDecls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcName := funcDecl.Name.Name
			if overwriteDecl := newFuncs[funcName]; overwriteDecl != nil {
				currentDecls[i] = overwriteDecl
				delete(newFuncs, funcName)
			}
		}
	}

	for _, newFuncDecl := range newFuncs {
		s.File.Decls = append(s.File.Decls, newFuncDecl)
	}

	return nil
}

func (s *Session) appendStatements(stmts ...ast.Stmt) {
	s.mainBody.List = append(s.mainBody.List, stmts...)
}

// Error is an exported error.
type Error string

// ErrContinue and ErrQuit are specific exported error types.
const (
	ErrContinue Error = "<continue input>"
	ErrQuit     Error = "<quit session>"
)

func (e Error) Error() string {
	return string(e)
}

func (s *Session) source(space bool) (string, error) {
	normalizeNodePos(s.mainFunc())

	var config *printer.Config
	if space {
		config = &printer.Config{
			Mode:     printer.UseSpaces,
			Tabwidth: 4,
		}
	} else {
		config = &printer.Config{
			Tabwidth: 8,
		}
	}

	var buf bytes.Buffer
	err := config.Fprint(&buf, s.Fset, s.File)
	return buf.String(), err
}

func (s *Session) reset() error {
	source, err := s.source(false)
	if err != nil {
		return err
	}

	file, err := parser.ParseFile(s.Fset, "gophernotes_session.go", source, parser.Mode(0))
	if err != nil {
		return err
	}

	s.File = file
	s.mainBody = s.mainFunc().Body

	return nil
}

// Eval handles the evaluation of code parsed from received messages
func (s *Session) Eval(in string) (string, bytes.Buffer, error) {
	debugf("eval >>> %q", in)

	s.clearQuickFix()
	s.storeMainBody()

	// Split the lines of the input to check for special commands.
	inLines := strings.Split(in, "\n")
	var nonImportLines []string
	inMultilineImport := false
	for idx, line := range inLines {
		// Extract non-special lines.
		trimLine := strings.TrimSpace(line)

		isImport := strings.HasPrefix(trimLine, "import")
		isSpecialCommand := strings.HasPrefix(line, ":")
		wasPartOfMultiline := inMultilineImport

		// If in a multi-line import and we find the closing paren, exit this mode
		if inMultilineImport && (strings.Count(line, "(")-strings.Count(line, ")")) < 0 {
			inMultilineImport = false
		}

		// If this import line has an opening paren and not a closing one, assume in multi-line import
		if isImport && (strings.Count(line, "(")-strings.Count(line, ")")) > 0 {
			inMultilineImport = true
		}

		if !isImport && !isSpecialCommand && !wasPartOfMultiline {
			nonImportLines = append(nonImportLines, line)
			continue
		}

		// Process special commands.
		var args []string
		for _, command := range commands {

			// Clear the args.
			args = []string{}

			// Extract any argument provided with the special command.
			arg := strings.TrimPrefix(line, ":"+command.name)
			if command.name == "import" {
				arg = strings.TrimPrefix(arg, "import")
			}
			switch {
			case arg == line:
				continue
			case strings.HasPrefix(strings.TrimSpace(arg), "("):
				advance := 1
				currentLine := inLines[idx+advance]
				for !strings.Contains(currentLine, ")") {
					args = append(args, currentLine)
					advance++
					currentLine = inLines[idx+advance]
				}
			default:
				args = append(args, arg)
			}

			// Apply the action associated with the special command.
			for _, arg = range args {
				if arg == "" || strings.HasPrefix(arg, " ") {
					arg = strings.TrimSpace(arg)
					_, err := command.action(s, arg)
					if err != nil {
						if err == ErrQuit {
							return "", bytes.Buffer{}, err
						}
						errorf("%s: %s", command.name, err.Error())
					}
				}
			}
		}
	}

	// Join the non-special lines back together for evaluation.
	in = strings.Join(nonImportLines, "\n")
	if len(in) == 0 {
		s.doQuickFix()
		return "", bytes.Buffer{}, nil
	}

	// Extract statements.
	priorListLength := len(s.mainBody.List)

	// Reenable stdout for the most recent execution
	restoreStdout := &ast.AssignStmt{
		Lhs:    []ast.Expr{ast.NewIdent("os.Stdout")},
		TokPos: token.NoPos,
		Tok:    token.ASSIGN,
		Rhs:    []ast.Expr{ast.NewIdent(oldStdout)},
	}
	s.appendStatements(restoreStdout)

	if err := s.separateEvalStmt(in); err != nil {
		// Remove any newly added statements as there was an error parsing
		// the cell
		s.mainBody.List = s.mainBody.List[0:priorListLength]
		return "", *bytes.NewBuffer([]byte(err.Error())), err
	}

	s.doQuickFix()

	output, stderr, runErr := s.Run()
	if runErr != nil || stderr.String() != "" {
		/*if exitErr, ok := runErr.(*exec.ExitError); ok {
			// if failed with status 2, remove the last statement
			if st, ok := exitErr.ProcessState.Sys().(syscall.WaitStatus); ok {
				if st.ExitStatus() == 2 {
					debugf("got exit status 2, popping out last input")
					s.restoreMainBody()
					runErr = nil
				}
			}
		}*/

		// Remove any newly added statements as there was an error executing
		// the cell
		s.mainBody.List = s.mainBody.List[0:priorListLength]
		return string(output), stderr, runErr
	}

	// Delete the single stdout redirection injected before the current execution so that it doesn't happen
	// for the next execution.
	s.mainBody.List = append(s.mainBody.List[:priorListLength], s.mainBody.List[priorListLength+1:]...)

	f, err := os.Create(s.FilePath)
	if err != nil {
		return string(output), stderr, err
	}
	err = printer.Fprint(f, s.Fset, s.File)
	if err != nil {
		return string(output), stderr, err
	}

	// Catch any unexpected stderr.
	if stderr.String() != "" {
		runErr = errors.New("Unexpected stderr from execution")
	}

	return string(output), stderr, runErr
}

func scopeDepthChange(line string) int {
	return strings.Count(line, "{") - strings.Count(line, "}")
}

// separateEvalStmt separates what can be evaluated via evalExpr from what cannot.
func (s *Session) separateEvalStmt(in string) error {
	var stmtLines []string
	var bracketCount int

	inLines := strings.Split(in, "\n")
	lastLineIdx := len(inLines) - 1

	for idx, line := range inLines {
		// If at the top level, try to evaluate the line as an expression
		if bracketCount == 0 {
			_, err := s.evalExpr(line)

			// Successfully parsed the single line as an expression so this line
			// is handled, continue to the next one
			if err == nil {
				continue
			}
		}

		// Adjust the scope depth by counting the opening and closing brackets on this
		// next line.
		bracketCount += scopeDepthChange(line)

		// Append the current line to the list of lines making up the multi line statement
		stmtLines = append(stmtLines, line)

		// If at the top level and the brackets are all matched up, put together the
		// multiline statement and evalualte it.
		// Note: this could be a single line if it contains all the matching braces but
		// is at the top level
		if bracketCount == 0 && len(stmtLines) > 0 {
			// Try and evaluate the statement. It fails then this is actually a
			// syntax error.
			if err := s.evalStmt(strings.Join(stmtLines, "\n"), idx != lastLineIdx); err != nil {
				return err
			}
			stmtLines = []string{}
		}
	}

	return nil
}

// cleanEvalStmt cleans up prior print statements etc.
func (s *Session) cleanEvalStmt(in string) error {
	var stmtLines []string

	inLines := strings.Split(in, "\n")

	for _, line := range inLines {
		beforeLines := len(s.mainBody.List)
		if expr, err := s.evalExpr(line); err == nil {
			// Remove the statement appended by evalExpr
			s.mainBody.List = s.mainBody.List[0:beforeLines]
			// If the statement is pure it can be skipped during the replay
			if s.isPureExpr(expr) {
				continue
			}
		}

		stmtLines = append(stmtLines, line)
	}

	if len(stmtLines) != 0 {
		if err := s.evalStmt(strings.Join(stmtLines, "\n"), true); err != nil {
			return err
		}
	}

	return nil
}

// storeMainBody stores current state of code so that it can be restored
// actually it saves the length of statements inside main()
func (s *Session) storeMainBody() {
	s.storedBodyLength = len(s.mainBody.List)
}

func (s *Session) restoreMainBody() {
	s.mainBody.List = s.mainBody.List[0:s.storedBodyLength]
}

// includeFiles imports packages and funcsions from multiple golang source
func (s *Session) includeFiles(files []string) {
	for _, file := range files {
		s.includeFile(file)
	}
}

func (s *Session) includeFile(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		errorf("%s", err.Error())
		return
	}

	if err = s.importPackages(content); err != nil {
		errorf("%s", err.Error())
		return
	}

	if err = s.importFile(content); err != nil {
		errorf("%s", err.Error())
	}

	infof("added file %s", file)
}

// importPackages includes packages defined on external file into main file
func (s *Session) importPackages(src []byte) error {
	astf, err := parser.ParseFile(s.Fset, "", src, parser.Mode(0))
	if err != nil {
		return err
	}

	for _, imt := range astf.Imports {
		debugf("import package: %s", imt.Path.Value)
		actionImport(s, imt.Path.Value)
	}

	return nil
}

// importFile adds external golang file to goRun target to use its function
func (s *Session) importFile(src []byte) error {
	// Don't need to same directory
	tmp, err := ioutil.TempFile(filepath.Dir(s.FilePath), "gore_external_")
	if err != nil {
		return err
	}

	ext := tmp.Name() + ".go"

	f, err := parser.ParseFile(s.Fset, ext, src, parser.Mode(0))
	if err != nil {
		return err
	}

	// rewrite to package main
	f.Name.Name = "main"

	// remove func main()
	for i, decl := range f.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if isNamedIdent(funcDecl.Name, "main") {
				f.Decls = append(f.Decls[0:i], f.Decls[i+1:]...)
				// main() removed from this file, we may have to
				// remove some unsed import's
				quickfix.QuickFix(s.Fset, []*ast.File{f})
				break
			}
		}
	}

	out, err := os.Create(ext)
	if err != nil {
		return err
	}
	defer out.Close()

	err = printer.Fprint(out, s.Fset, f)
	if err != nil {
		return err
	}

	debugf("import file: %s", ext)
	s.ExtraFilePaths = append(s.ExtraFilePaths, ext)
	s.ExtraFiles = append(s.ExtraFiles, f)

	return nil
}

// fixImports formats and adjusts imports for the current AST.
func (s *Session) fixImports() error {

	var buf bytes.Buffer
	err := printer.Fprint(&buf, s.Fset, s.File)
	if err != nil {
		return err
	}

	formatted, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return err
	}

	s.File, err = parser.ParseFile(s.Fset, "", formatted, parser.Mode(0))
	if err != nil {
		return err
	}
	s.mainBody = s.mainFunc().Body

	return nil
}

func (s *Session) includePackage(path string) error {
	pkg, err := build.Import(path, ".", 0)
	if err != nil {
		var err2 error
		pkg, err2 = build.ImportDir(path, 0)
		if err2 != nil {
			return err // return package path import error, not directory import error as build.Import can also import directories if "./foo" is specified
		}
	}

	files := make([]string, len(pkg.GoFiles))
	for i, f := range pkg.GoFiles {
		files[i] = filepath.Join(pkg.Dir, f)
	}
	s.includeFiles(files)

	return nil
}
