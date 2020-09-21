package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/plantuml"
	"github.com/xwb1989/sqlparser"
	"sort"
	"strconv"
	"strings"
)

func generateSQL(md *Markdown, bc *ddd.BoundedContextSpec, layer *ddd.MySQLLayerSpec) {
	md.H4("Persistence layer: *" + layer.Name() + "*")
	md.P("The *mysql* persistence layer for the domain *" + bc.Name() + "* consists of\n" +
		textMigrations(len(layer.Migrations())) + " and has been evolved as shown in the following table:",
	)
	md.TableHeader("Introduced at", "Description", "Statements")
	lastestDateTime := ""
	for _, migration := range layer.Migrations() {
		lastestDateTime = migration.DateTime()
		stats := map[string]int{}
		for _, statement := range migration.RawStatements() {
			sql, err := sqlparser.ParseStrictDDL(string(statement))
			if err != nil {
				panic("not yet validated")
			}
			switch t := sql.(type) {
			case *sqlparser.DDL:
				stats[t.Action] = stats[t.Action] + 1
			case *sqlparser.Insert:
				stats["insert"] = stats["insert"] + 1
			case *sqlparser.Update:
				stats["update"] = stats["update"] + 1
			case *sqlparser.Delete:
				stats["delete"] = stats["delete"] + 1
			default:
				stats["other"] = stats["other"] + 1
			}
		}

		var sortedKeys []string
		for k, _ := range stats {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		tmp := &strings.Builder{}
		for _, key := range sortedKeys {
			tmp.WriteString(strconv.Itoa(stats[key]))
			tmp.WriteString("x ")
			tmp.WriteString(key)
			tmp.WriteString("<br>")
		}
		md.TableRow(migration.DateTime(), migration.Description(), tmp.String())
	}

	md.P("After applying all migrations (as of " + lastestDateTime + "), the final sql schema looks like follows:")

	diag := md.UML(layer.Name() + "-er")
	db := buildActualSchema(layer)
	for _, table := range db.tables {
		dTab := plantuml.NewClass(table.name)
		for _, col := range table.columns {
			dTab.AddAttrs(plantuml.Attr{
				Visibility: plantuml.Public,
				Abstract:   false,
				Static:     false,
				Name:       col.name,
				Type:       col.kind,
			})
		}
		diag.Add(dTab)
	}

}

type sqlDatabase struct {
	tables []*sqlTable
}

type sqlTable struct {
	name    string
	columns []*sqlCol
}

type sqlCol struct {
	name string
	kind string
}

func buildActualSchema(layer ddd.SQLLayer) *sqlDatabase {
	db := &sqlDatabase{}
	for _, migration := range layer.Migrations() {
		for _, statement := range migration.RawStatements() {
			sql, err := sqlparser.ParseStrictDDL(string(statement))
			if err != nil {
				panic("not yet validated")
			}

			switch t := sql.(type) {
			case *sqlparser.DDL:
				switch t.Action {
				case sqlparser.CreateStr:
					table := &sqlTable{
						name: t.NewName.Name.String(),
					}
					for _, column := range t.TableSpec.Columns {
						col := &sqlCol{
							name: column.Name.String(),
							kind: column.Type.DescribeType(),
						}
						table.columns = append(table.columns, col)
					}
					db.tables = append(db.tables, table)
				}
			}
		}
	}
	return db
}

func textMigrations(q int) string {
	if q == 1 {
		return "one migration"
	}

	return strconv.Itoa(q) + " migrations"
}
