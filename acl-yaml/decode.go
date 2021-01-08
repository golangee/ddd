package aclyaml

import (
	"github.com/golangee/architecture/acl"
)

type validateable interface {
	Validate() error
}

// Parse loads up all *.yaml files recursively from the given directory. The given directory name
// is considered the root package.
func Parse(dir string) (acl.Pkg, error) {
	return NewYamlPkg(dir)
}
