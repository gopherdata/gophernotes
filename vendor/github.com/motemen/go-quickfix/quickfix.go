// Package quickfix provides functions for fixing Go ASTs
// that are well typed but "go build" refuses to build.
package quickfix

import (
	"fmt"
	"regexp"
	"strings"

	"go/ast"
	"go/importer"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ast/astutil"
)

var (
	declaredNotUsed        = regexp.MustCompile(`^([a-zA-Z0-9_]+) declared but not used$`)
	importedNotUsed        = regexp.MustCompile(`^(".+") imported but not used`)
	noNewVariablesOnDefine = "no new variables on left side of :="
)

type Config struct {
	Fset     *token.FileSet
	Files    []*ast.File
	TypeInfo *types.Info
	MaxTries int
}

func QuickFix(fset *token.FileSet, files []*ast.File) (err error) {
	config := Config{
		Fset:     fset,
		Files:    files,
		MaxTries: 10,
	}
	return config.QuickFix()
}

// QuickFix rewrites AST files of same package so that they pass go build.
// For example:
//   v declared but not used             -> append `_ = v`
//   "p" imported but not used           -> rewrite to `import _ "p"`
//   no new variables on left side of := -> rewrite `:=` to `=`
//
// TODO implement hardMode, which removes errorneous code rather than adding
func (c Config) QuickFix() (err error) {
	maxTries := 10
	if c.MaxTries > 0 {
		maxTries = c.MaxTries
	}
	for i := 0; i < maxTries; i++ {
		var foundError bool
		foundError, err = c.QuickFixOnce()
		if !foundError {
			return nil
		}
	}

	return
}

type tracedVisitor struct {
	path  []ast.Node
	visit func(ast.Node, []ast.Node) bool
}

func (v tracedVisitor) Visit(node ast.Node) ast.Visitor {
	if v.visit(node, v.path) {
		return tracedVisitor{
			path:  append([]ast.Node{node}, v.path...),
			visit: v.visit,
		}
	}

	return nil
}

func traverseAST(node ast.Node, visit func(ast.Node, []ast.Node) bool) {
	v := tracedVisitor{
		visit: visit,
	}
	ast.Walk(v, node)
}

// pkgsWithSideEffect are set of packages which are known to provide APIs by
// blank identifier import (import _ "p").
var pkgsWithSideEffect = map[string]bool{}

func init() {
	for _, path := range []string{
		"expvar",
		"image/gif",
		"image/jpeg",
		"image/png",
		"net/http/pprof",
		"unsafe",
		"golang.org/x/image/bmp",
		"golang.org/x/image/tiff",
		"golang.org/x/image/vp8",
		"golang.org/x/image/vp81",
		"golang.org/x/image/webp",
		"golang.org/x/tools/go/gcimporter",
	} {
		pkgsWithSideEffect[`"`+path+`"`] = true
	}
}

func RevertQuickFix(fset *token.FileSet, files []*ast.File) error {
	config := Config{
		Fset:     fset,
		Files:    files,
		MaxTries: 10,
	}
	return config.RevertQuickFix()
}

// RevertQuickFix reverts possible quickfixes introduced by QuickFix.
// This may result to non-buildable source, and cannot reproduce the original
// code before prior QuickFix.
// For example:
//   `_ = v`        -> removed
//   `import _ "p"` -> rewritten to `import "p"`
func (c Config) RevertQuickFix() (err error) {
	fset := c.Fset
	files := c.Files

	nodeToRemove := map[ast.Node]bool{}

	for _, f := range files {
		ast.Inspect(f, func(node ast.Node) bool {
			if assign, ok := node.(*ast.AssignStmt); ok {
				if len(assign.Lhs) == 1 && isBlankIdent(assign.Lhs[0]) &&
					len(assign.Rhs) == 1 && isIdent(assign.Rhs[0]) {
					// The statement is `_ = v`
					nodeToRemove[node] = true
				}

				return false
			} else if imp, ok := node.(*ast.ImportSpec); ok {
				if isBlankIdent(imp.Name) && !pkgsWithSideEffect[imp.Path.Value] {
					// The spec is `import _ "p"` and p is not a package that
					// provides "side effects"
					imp.Name = nil
				}

				return false
			}

			return true
		})

		for len(nodeToRemove) > 0 {
			traverseAST(f, func(node ast.Node, nodepath []ast.Node) bool {
				if nodeToRemove[node] {
					parent := nodepath[0]
					if removeChildNode(node, parent) == false {
						err = fmt.Errorf(
							"BUG: could not remove node: %s (in: %s)",
							fset.Position(node.Pos()),
							fset.Position(parent.Pos()),
						)
					}
					delete(nodeToRemove, node)
					return false
				}

				return true
			})
		}
	}

	return
}

func (c Config) QuickFixOnce() (bool, error) {
	fset := c.Fset
	files := c.Files

	errs := []error{}
	config := &types.Config{
		Error: func(err error) {
			errs = append(errs, err)
		},
		Importer: importer.Default(),
	}

	_, err := config.Check("_quickfix", fset, files, c.TypeInfo)
	if err == nil {
		return false, nil
	}

	// apply fixes on AST later so that we won't break funcs that inspect AST by positions
	fixes := map[error]func() bool{}
	unhandled := ErrorList{}

	foundError := len(errs) > 0

	for _, err := range errs {
		err, ok := err.(types.Error)
		if !ok {
			unhandled = append(unhandled, err)
			continue
		}

		f := findFile(c.Files, err.Pos)
		if f == nil {
			e := ErrCouldNotLocate{
				Err:  err,
				Fset: fset,
			}
			unhandled = append(unhandled, e)
			continue
		}

		nodepath, _ := astutil.PathEnclosingInterval(f, err.Pos, err.Pos)

		var fix func() bool

		// - "%s declared but not used"
		// - "%q imported but not used" (+ " as %s")
		// - "label %s declared but not used" TODO
		// - "no new variables on left side of :="
		if m := declaredNotUsed.FindStringSubmatch(err.Msg); m != nil {
			identName := m[1]
			fix = func() bool {
				return fixDeclaredNotUsed(nodepath, identName)
			}
		} else if m := importedNotUsed.FindStringSubmatch(err.Msg); m != nil {
			pkgPath := m[1] // quoted string, but it's okay because this will be compared to ast.BasicLit.Value.
			fix = func() bool {
				return fixImportedNotUsed(nodepath, pkgPath)
			}
		} else if err.Msg == noNewVariablesOnDefine {
			fix = func() bool {
				return fixNoNewVariables(nodepath)
			}
		} else {
			unhandled = append(unhandled, err)
		}

		if fix != nil {
			fixes[err] = fix
		}
	}

	for err, fix := range fixes {
		if fix() == false {
			unhandled = append(unhandled, err)
		}
	}

	return foundError, unhandled.any()
}

func fixDeclaredNotUsed(nodepath []ast.Node, identName string) bool {
	// insert "_ = x" to supress "declared but not used" error
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("_")},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{ast.NewIdent(identName)},
	}
	return appendStmt(nodepath, stmt)
}

func fixImportedNotUsed(nodepath []ast.Node, pkgPath string) bool {
	for _, node := range nodepath {
		if f, ok := node.(*ast.File); ok {
			for _, imp := range f.Imports {
				if imp.Path.Value == pkgPath {
					// make this import spec anonymous one
					imp.Name = ast.NewIdent("_")
					return true
				}
			}
		}
	}
	return false
}

func fixNoNewVariables(nodepath []ast.Node) bool {
	for _, node := range nodepath {
		switch node := node.(type) {
		case *ast.AssignStmt:
			if node.Tok == token.DEFINE {
				node.Tok = token.ASSIGN
				return true
			}

		case *ast.RangeStmt:
			if node.Tok == token.DEFINE {
				node.Tok = token.ASSIGN
				return true
			}
		}
	}
	return false
}

type ErrorList []error

func (errs ErrorList) any() error {
	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (errs ErrorList) Error() string {
	s := []string{fmt.Sprintf("%d error(s):", len(errs))}
	for _, e := range errs {
		s = append(s, fmt.Sprintf("- %s", e))
	}
	return strings.Join(s, "\n")
}

func appendStmt(nodepath []ast.Node, stmt ast.Stmt) bool {
	for _, node := range nodepath {
		switch node := node.(type) {
		case *ast.BlockStmt:
			if node.List == nil {
				node.List = []ast.Stmt{}
			}
			node.List = append(node.List, stmt)

		case *ast.CaseClause:
			if node.Body == nil {
				node.Body = []ast.Stmt{}
			}
			node.Body = append(node.Body, stmt)

		case *ast.CommClause:
			if node.Body == nil {
				node.Body = []ast.Stmt{}
			}
			node.Body = append(node.Body, stmt)

		case *ast.RangeStmt:
			if node.Body == nil {
				node.Body = &ast.BlockStmt{}
			}
			if node.Body.List == nil {
				node.Body.List = []ast.Stmt{}
			}
			node.Body.List = append(node.Body.List, stmt)

		default:
			continue
		}

		return true
	}

	return false
}

func removeChildNode(child, parent ast.Node) bool {
	switch parent := parent.(type) {
	case *ast.BlockStmt:
		removeFromStmtList(child, parent.List)
		return true
	case *ast.CaseClause:
		removeFromStmtList(child, parent.Body)
		return true
	case *ast.CommClause:
		removeFromStmtList(child, parent.Body)
		return true
	case *ast.RangeStmt:
		removeFromStmtList(child, parent.Body.List)
		return true
	}

	return false
}

// removeFromStmtList remove node from slice of statements list. This function
// modifies list in-place and pads rest of the slice with ast.EmptyStmt.
func removeFromStmtList(node ast.Node, list []ast.Stmt) bool {
	for i, s := range list {
		if s == node {
			for ; i < len(list)-1; i++ {
				list[i] = list[i+1]
			}
			list[len(list)-1] = &ast.EmptyStmt{}
			return true
		}
	}

	return false
}

func findFile(files []*ast.File, pos token.Pos) *ast.File {
	for _, f := range files {
		if f.Pos() <= pos && pos < f.End() {
			return f
		}
	}

	return nil
}

func isIdent(node ast.Node) bool {
	if node == nil {
		return false
	}

	_, ok := node.(*ast.Ident)
	return ok
}

func isBlankIdent(node ast.Node) bool {
	if node == nil {
		return false
	}

	ident, ok := node.(*ast.Ident)
	return ok && ident != nil && ident.Name == "_"
}

type ErrCouldNotLocate struct {
	Err  types.Error
	Fset *token.FileSet
}

func (e ErrCouldNotLocate) Error() string {
	return fmt.Sprintf("cannot find file for error %q: %s (%d)", e.Err.Error(), e.Fset.Position(e.Err.Pos), e.Err.Pos)
}
