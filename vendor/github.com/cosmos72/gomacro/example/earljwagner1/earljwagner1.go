package main

import (
	"fmt"
	"reflect"

	"github.com/cosmos72/gomacro/fast"
	"github.com/cosmos72/gomacro/imports"
)

// example from Earl J Wagner use case:
// interactively replace a (bugged) compiled function with a (corrected) interpreted one
func main() {
	// 1. create the fast interpreter.
	ir := fast.New()

	// 2. tell the interpreter about our compiled function Cube() in package "github.com/cosmos72/gomacro/example/earljwagner1"
	// An alternative solution is to run the interpreter interactively, and at its REPL enter the command:
	// import _i "package/to/generate/imports/for"
	// (note: the _i is fundamental)
	// This will create a file x_package.go in the imported package - just recompile and rerun you program:
	// the interpreter will now be able to 'import "package/to/generate/imports/for"'
	// without the need to dynamically compile and load a plugin
	imports.Packages["github.com/cosmos72/gomacro/example/earljwagner1"] = imports.Package{
		Binds: map[string]reflect.Value{
			"Cube": reflect.ValueOf(Cube),
		},
	}

	// 3. tell the interpreter to import the package containing our Cube()
	//    alternatively, one could use
	//    ir.ImportPackage("earljwagner1", "github.com/cosmos72/gomacro/example/earljwagner1")
	ir.Eval(`import "github.com/cosmos72/gomacro/example/earljwagner1"`)

	// 4. execute compiled Cube() - and realise it's bugged
	xcube, _ := ir.Eval1("earljwagner1.Cube(3.0)")
	fmt.Printf("compiled earljwagner1.Cube(3.0) = %f\n", xcube.Interface().(float64))

	// 5. tell the interpreter to switch to package "github.com/cosmos72/gomacro/example/earljwagner1"
	//    alternatively, at REPL one would type (note the quotes):
	//      package "github.com/cosmos72/gomacro/example/earljwagner1"
	ir.ChangePackage("earljwagner1", "github.com/cosmos72/gomacro/example/earljwagner1")

	// 6. the compiled function Cube() can now be invoked without package prefix
	xcube, _ = ir.Eval1("Cube(4.0)")
	fmt.Printf("compiled Cube(4.0) = %f\n", xcube.Interface().(float64))

	// 7. define an interpreted function Cube(), replacing the compiled one
	ir.Eval("func Cube(x float64) float64 { return x*x*x }")

	// 8. invoke the interpreted function Cube() - the bug is solved :)
	xcube, _ = ir.Eval1("Cube(4.0)")
	fmt.Printf("interpreted Cube(4.0) = %f\n", xcube.Interface().(float64))

	// 9. note: compiled code will *NOT* automatically know about the bug-fixed Cube() living inside the interpreter.
	//    One solution is to stay inside the interpreter REPL and use interpreted functions.
	//    Another solution is to extract the bug-fixed function from the interpreter and use it,
	//    for example by storing it inside imports.Packages
	imports.Packages["github.com/cosmos72/gomacro/example/earljwagner1"].Binds["Cube"] = ir.ValueOf("Cube")
}

func Cube(x float64) float64 {
	return x*x*x - 1 // intentionally bugged
}
