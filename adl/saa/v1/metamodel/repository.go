package metamodel

import (
	"github.com/golangee/architecture/adl/saa/v1/spec"
)

// A Repository declares a repository and methods with data bindings.
type Repository struct {
	Name    model.Identifier
	Methods []Method
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context.
type Method struct {
	Name    model.Identifier
	Query   model.StrLit
	Prepare []model.Identifier
	Map     []model.Identifier
}
