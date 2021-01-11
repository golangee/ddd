package modlet

import "github.com/golangee/architecture/objn"

type Project struct {
}

type Modlet interface {
	Apply(prj Project, node objn.Node) error
}
