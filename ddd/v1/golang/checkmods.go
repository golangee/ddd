package golang

import (
	"fmt"
	"github.com/golangee/reflectplus/mod"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (g *generator) loadTargetModules() error {
	if g.opts.ArchDir == "" {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("no working directory: %err", err)
		}

		mods, err := mod.List(dir)
		if err != nil {
			return fmt.Errorf("'%s': cannot auto detect current module context: %w", dir, err)
		}

		g.arch = mods
		g.opts.ArchDir = mods.Main().Dir
		log.Printf("auto detected architecture project '%s' at '%s'\n", mods.Main().Path, mods.Main().Dir)
	}

	if strings.HasPrefix(g.opts.ServerDir, "../") {
		g.opts.ServerDir = filepath.Clean(filepath.Join(g.opts.ArchDir, g.opts.ServerDir))
	}

	mods, err := mod.List(g.opts.ServerDir)
	if err != nil {
		return fmt.Errorf("server '%s' is invalid: %w", g.opts.ServerDir, err)
	}

	g.server = mods
	log.Printf("server '%s' at '%s'\n", mods.Main().Path, mods.Main().Dir)

	if strings.HasPrefix(g.opts.ClientDir, "../") {
		g.opts.ClientDir = filepath.Clean(filepath.Join(g.opts.ArchDir, g.opts.ClientDir))
	}

	mods, err = mod.List(g.opts.ClientDir)
	if err != nil {
		return fmt.Errorf("client '%s' is invalid: %w", g.opts.ServerDir, err)
	}

	g.client = mods
	log.Printf("client '%s' at '%s'\n", mods.Main().Path, mods.Main().Dir)

	if g.arch.Main().Path == g.server.Main().Path {
		log.Printf("WARN: architecture and server share the same module\n")
	}

	if g.arch.Main().Path == g.client.Main().Path {
		log.Printf("WARN: architecture and client share the same module\n")
	}

	return nil
}
