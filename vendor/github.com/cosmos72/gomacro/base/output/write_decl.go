package output

import (
	"fmt"
	"go/ast"
	"io"
)

func (o *Output) WriteDeclsToStream(out io.Writer, packagePath string,
	imports []*ast.GenDecl, declarations []ast.Decl, statements []ast.Stmt) {

	fmt.Fprintf(out, "package %s\n\n", packagePath)

	for _, imp := range imports {
		fmt.Fprintln(out, o.toPrintable("%v", imp))
	}
	if len(imports) != 0 {
		fmt.Fprintln(out)
	}
	for _, decl := range declarations {
		fmt.Fprintln(out, o.toPrintable("%v", decl))
	}
	if len(statements) != 0 {
		fmt.Fprint(out, "\nfunc init() {\n")
		config.Indent = 1
		defer func() {
			config.Indent = 0
		}()
		for _, stmt := range statements {
			fmt.Fprintln(out, o.toPrintable("%v", stmt))
		}
		fmt.Fprint(out, "}\n")
	}
}
