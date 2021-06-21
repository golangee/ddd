package golang

import (
	"github.com/golangee/architecture/arc/adl"
)

// normalizeNames replaces special variable names within various places.
//  $MOD => the full qualified module path (e.g. github.com/mycompany/mymodule)
//  $BC => the full qualified bounded context path (e.g. github.com/mycompany/mymodule/internal/mydomain/bc)
func normalizeNames(src *adl.Module) {
	mod := src.Generator.Go.Module.String()
	ctx := adl.Ctx{
		Mod: mod,
	}

	for _, executable := range src.Executables {
		executable.Normalize(ctx)
	}

	for _, bc := range src.BoundedContexts {
		bc.Normalize(ctx)
	}
}
