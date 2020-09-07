package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
)

// Generate takes the ddd model and writes a README.md including graphics into the given directory.
func Generate(targetDir string, app *ddd.AppSpec) error {
	ctx := &genctx{
		spec: app,
	}

	if err := generateDocument(ctx); err != nil {
		return err
	}

	return ctx.emit(targetDir)
}
