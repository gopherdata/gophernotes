package genimport

import (
	"fmt"
	"go/build"
	"os"
	"strings"
)

const sep = string(os.PathSeparator)

// GoGenerateMain allows gomacro to be run under
// go generate. It is used to write new x_package.go
// import bindings for a package. Thus `go generate`
// will automatically update your bindings.
//
// To use, add a comment to one go file in your package:
//
//`//go:generate gomacro -g .`
//
// to import the current dir; or one like
//
//`//go:generate gomacro -g github.com/cosmos72/gomacro/classic`
//
// to specify the exact import path. The second, specific
// form, may be necessary if we cannot detect the GOPATH
// environment variable.
//
func GoGenerateMain(arg []string, imp *Importer) error {
	var pkgpath string
	narg := len(arg)
	switch {
	case narg == 0 || (narg > 0 && arg[0] == "."):
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("gomacro -g: error getting current dir: %v", err)
		}
		gopath := build.Default.GOPATH
		prefix := gopath + sep + "src" + sep
		if strings.HasPrefix(cwd, prefix) {
			pkgpath = cwd[len(prefix):]
		} else {
			// guess it is after the first `src` in cwd,
			// since traditionally all packages are
			// after $GOPATH/src/
			splt := strings.SplitN(cwd, sep+"src"+sep, 2)
			if len(splt) <= 1 {
				return fmt.Errorf("gomacro -g: unable to detect current package, please specify it")
			}
			pkgpath = splt[1]
		}
	default:
		pkgpath = arg[0]
	}
	_, err := imp.ImportPackageOrError("_i", pkgpath)
	return err
}
