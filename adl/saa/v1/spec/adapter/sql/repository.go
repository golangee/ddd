package sql

import (
	"github.com/golangee/architecture/adl/saa/v1/spec"
)

// A Repository declares a repository and methods with data bindings.
type Repository struct {
	Name    spec.Identifier
	Methods []Method
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context.
type Method struct {
	Name    spec.Identifier
	Query   spec.StrLit
	Prepare []spec.Identifier
	Map     []spec.Identifier
}
