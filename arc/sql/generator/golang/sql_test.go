package golang

import (
	"embed"
	"fmt"
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/golang"
	"github.com/golangee/src/stdlib"
	"testing"
)

//go:embed *.sql
var fs embed.FS

func TestRenderRepository(t *testing.T) {
	prj := createProject(t)

	ctx := createCtx(t)
	if err := RenderSQL(prj, ctx); err != nil {
		t.Fatal(token.Explain(err))
	}

	renderer := golang.NewRenderer(golang.Options{})
	a, err := renderer.Render(prj)
	if err != nil {
		fmt.Println(a)
		t.Fatal(err)
	}

	fmt.Println(a)
}

func lits(lit ...string) []token.String {
	var r []token.String
	for _, s := range lit {
		r = append(r, token.NewString(s))
	}
	return r
}

func createCtx(t *testing.T) *sql.Ctx {
	t.Helper()

	mod := token.NewString("github.com/worldiety/supportiety")
	pkg := token.NewString("github.com/worldiety/supportiety/tickets/core")

	return &sql.Ctx{
		Dialect:    sql.MySQL,
		Mod:        mod,
		Pkg:        pkg,
		Migrations: createMigrations(t),
		Repositories: []sql.Repository{
			{
				Implements: token.NewString("github.com/worldiety/supportiety/tickets/core.TicketRepository"),
				Methods: []sql.Method{
					{
						Name:    token.NewString("CreateTicket"),
						Query:   token.NewString("INSERT INTO tickets VALUES (?)"),
						Mapping: sql.ExecOne{In: lits("id")},
					},

					{
						Name:  token.NewString("CreateManyTickets"),
						Query: token.NewString("INSERT INTO tickets VALUES (?)"),
						Mapping: sql.ExecMany{
							Slice: token.NewString("ids"),
							In:    lits("ids[i]"),
						},
					},

					{
						Name:  token.NewString("FindAll"),
						Query: token.NewString("SELECT * FROM tickets"),
						Mapping: sql.QueryMany{
							In:  nil,
							Out: lits(".ID", ".Name", ".Desc"),
						},
					},

					{
						Name:  token.NewString("Count"),
						Query: token.NewString("SELECT COUNT(*) FROM tickets"),
						Mapping: sql.QueryOne{
							In:  nil,
							Out: lits("."),
						},
					},

					{
						Name:    token.NewString("DeleteTicket"),
						Query:   token.NewString("DELETE FROM tickets where id=?"),
						Mapping: sql.ExecOne{In: lits("id")},
					},

					{
						Name:  token.NewString("FindTicket"),
						Query: token.NewString("SELECT * FROM tickets where id=?"),
						Mapping: sql.QueryOne{
							In:  lits("id"),
							Out: lits(".ID", ".Name"),
						},
					},
				},
			},
			{
				Implements: token.NewString("github.com/worldiety/supportiety/tickets/core.TicketFiles"),
				Methods: []sql.Method{

				},
			},
		},
	}
}

func createProject(t *testing.T) *ast.Prj {
	prj := ast.NewPrj("test")
	prj.AddModules(
		ast.NewMod("github.com/worldiety/supportiety").
			AddPackages(
				ast.NewPkg("github.com/worldiety/supportiety/tickets/core").
					AddFiles(
						ast.NewFile("repos.go").
							AddNodes(
								ast.NewStruct("Ticket").
									SetComment("...represents a domain ticket entity"),
								ast.NewInterface("TicketRepository").
									SetComment("...provides CRUD access to Tickets.").
									AddMethods(
										ast.NewFunc("CreateTicket").
											SetComment("...creates a Ticket.").
											AddParams(
												ast.NewParam("id", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

										ast.NewFunc("CreateManyTickets").
											SetComment("...creates a Ticket.").
											AddParams(
												ast.NewParam("ids", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl(stdlib.UUID))),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

										ast.NewFunc("FindTicket").
											SetComment("...find a Ticket.").
											AddParams(
												ast.NewParam("id", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl("Ticket")),
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

										ast.NewFunc("DeleteTicket").
											SetComment("...deletes a Ticket.").
											AddParams(
												ast.NewParam("id", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

										ast.NewFunc("FindAll").
											SetComment("...find all Tickets.").
											AddParams(
												ast.NewParam("id", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl("Ticket"))),
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

										ast.NewFunc("Count").
											SetComment("...counts all Tickets.").
											AddParams(
												ast.NewParam("id", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Int64)),
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),
									),


								ast.NewInterface("TicketFiles").
									SetComment("...connects files and tickets.").
									AddMethods(
										ast.NewFunc("AttachFile").
											SetComment("...connects a file and a ticket.").
											SetComment("...creates a Ticket.").
											AddParams(
												ast.NewParam("ticketId", ast.NewSimpleTypeDecl(stdlib.UUID)),
												ast.NewParam("fileId", ast.NewSimpleTypeDecl(stdlib.UUID)),
											).
											AddResults(
												ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
											),

									),
							),
					),
			),
	)

	return prj
}

func createMigrations(t *testing.T) []*sql.Migration {
	t.Helper()

	entries, err := fs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	var migrations []*sql.Migration
	for i, entry := range entries {
		ts, name, err := sql.ParseMigrationName(entry.Name())
		if err != nil {
			t.Fatal(err)
		}

		file, err := fs.Open(entry.Name())
		if err != nil {
			t.Fatal(err)
		}

		stmts, err := sql.ParseStatements(file)
		if err != nil {
			t.Fatal(err)
		}

		strName := token.NewString(name)
		strName.BeginPos.File = "sql_test.go"
		strName.BeginPos.Line = i + 1

		migrations = append(migrations, &sql.Migration{
			ID:         ts,
			Name:       strName,
			Statements: stmts,
		})
	}

	return migrations
}
