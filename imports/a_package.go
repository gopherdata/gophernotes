package imports

import (
	"github.com/cosmos72/gomacro/base/paths"
	i "github.com/cosmos72/gomacro/imports"
)

type PackageUnderlying = i.PackageUnderlying

type Package = i.Package

type PackageMap = i.PackageMap

// i.Packages is a map, copying it only makes a reference
// => Packages and i.Packages are kept in sync automatically
var Packages = i.Packages

// used to locate this package path via reflection
type findMe struct{}

func init() {
	// tell gomacro where to write import files
	paths.SymbolFromImportsPackage = findMe{}
}
