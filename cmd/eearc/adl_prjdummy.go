package eearc

import (
	"encoding/json"
	. "github.com/golangee/architecture/arc/adl"
)

func createWorkspace() *Project {
	return NewProject("supportiety").
		AddProjects(
			NewModule("server").
				SetGenerator(
					NewGenerator().
						SetOutDir("../../testdata/workspace/server").
						SetGo(NewGolang().SetModName("github.com/golangee/architecture/testdata/workspace/server")),
				).
				SetDomain(
					NewDomain("github.com/golangee/architecture/testdata/workspace/server/tickets").
						SetCore(
							NewPackage().AddRepositories(
								NewInterface("Tickets", "...provides CRUD access to Tickets.").
									AddMethods(
										NewMethod("CreateTicket", "...creates a Ticket.").
											AddIn("id", "...the unique ticket id", "uuid!").
											AddOut("", "...if anything goes wrong", "error!"),
									),
							),
						),
				),
		)
}

func toJson(i interface{}) string {
	buf, err := json.MarshalIndent(i, " ", " ")
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func wsFromJson(buf string) *Project {
	tmp := &Project{}
	if err := json.Unmarshal([]byte(buf), tmp); err != nil {
		panic(err)
	}

	return tmp
}
