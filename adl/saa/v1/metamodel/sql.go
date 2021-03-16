package metamodel

import (
	"github.com/golangee/architecture/adl/saa/v1/spec"
	"sort"
	"time"
)

// Dialect determines the kind of sql dialect.
type Dialect string

const (
	MySQL Dialect = "mysql"
)

// A Migration represents a transactional group of sql migration statements. All of them should be applied or none.
// However due to SQL nature, many engines do not support that well with CREATE/DROP TABLE statements.
type Migration struct {
	ID         time.Time
	Statements []model.StrLit
}

// Migrations is a slice of Migration.
type Migrations []Migration

// NewMigrations takes ownership of m and sorts the migrations ascending.
func NewMigrations(m ...Migration) Migrations {
	var r Migrations
	r = m
	sort.Sort(r)

	return r
}

func (m Migrations) Len() int {
	return len(m)
}

func (m Migrations) Less(i, j int) bool {
	return m[i].ID.Unix() < m[i].ID.Unix()
}

func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
