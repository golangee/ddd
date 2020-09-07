package architecture

import (
	"log"
	"os"
	"path/filepath"
)

type Project struct {
	// Dir is the root of the currently executing context, which is usually the compiled go architecture program.
	Dir string
}

// Detect returns a project from the current context. This contains either the modules root, if executed from within
// an IDE or makefile or just the current working directory otherwise.
func Detect() (*Project, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rootDir := cwd
	for len(cwd) > 1 {
		if isMod(cwd) {
			rootDir = cwd
			break
		}

		cwd = filepath.Dir(cwd)
	}

	log.Printf("detected architecture base directory: %s\n", rootDir)

	return &Project{Dir: rootDir}, nil
}

// File returns the absolute file path for the current project.
func (p *Project) File(name string) string {
	return filepath.Clean(filepath.Join(p.Dir, name))
}

func isMod(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		return false
	}

	return true
}
