package modlet

import "github.com/golangee/architecture/yast"

type Project struct {
}

type Modlet interface {
	Apply(prj Project, node yast.Node) error
}
