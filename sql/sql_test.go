package sql_test

import (
	. "github.com/golangee/ddd/sql"
	"testing"
)

func TestMigrate(t *testing.T) {
	Migrations(
		Migrate("2006-01-02T15:04:05",
			CreateTable("users"),
		),
	)
}
