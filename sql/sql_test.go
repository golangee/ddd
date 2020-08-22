package sql_test

import (
	. "github.com/golangee/ddd/sql"
	"testing"
)

func TestMigrate(t *testing.T) {
	Migrations(
		Migrate("2006-01-02T15:04:05",
			CreateTable("users",
				Columns(
					Int("id", 11, AutoIncrement()),
					Varchar("name", 255, NotNull()),
					Binary("uuid", 16),
				),
				PrimaryKey("id", "name"),
				ForeignKey("uuid", "objects", "id"),
			),
			AlterTable("users",
				AddColumns(
					Int("num", 3, NotNull()),
				),
				DropColumns("name"),
			),
		),
		Migrate("2006-01-02T15:04:05",
			Statement("CREATE TABLE ..."),
		),
	)
}
