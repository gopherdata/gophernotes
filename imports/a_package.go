package imports

import (
	"github.com/cosmos72/gomacro/base/paths"
	"github.com/cosmos72/gomacro/imports"
)

type PackageUnderlying = imports.PackageUnderlying

type Package = imports.Package

type PackageMap = imports.PackageMap

// imports.Packages is a map, copying it only makes a reference
// => Packages and imports.Packages are kept in sync automatically
var Packages = imports.Packages

// used to locate this package path via reflection
type findMe struct{}

func init() {
	// tell gomacro where to write import files
	paths.SymbolFromImportsPackage = findMe{}
}
