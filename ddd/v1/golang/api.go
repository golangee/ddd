package golang

import "github.com/golangee/architecture/ddd/v1"

// Options of the generator
type Options struct {
	// ArchDir is where the architecture project is (usually empty for auto detection).
	// Auto detection will set this to the module root, of the module which calls Generate.
	ArchDir string

	// ServerDir is a path to an existing go module. If it starts with a ../ it is resolved
	// to the ArchDir.
	ServerDir string

	// ClientDir is a path to an existing go module. If it starts with a ../ it is resolved
	// to the ArchDir.
	ClientDir string

	// DoNotClean if false, every file with the according "generated header" will be
	// removed before regeneration. This ensures that if the architecture changes,
	// obsolete files are automatically removed.
	DoNotClean bool
}

// Generate takes the ddd model and writes it into the filesystem.
func Generate(opts Options, app *ddd.AppSpec) error {
	gen := newGenerator(opts, app)
	return gen.make()
}
