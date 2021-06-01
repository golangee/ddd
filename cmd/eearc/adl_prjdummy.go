package eearc

import (
	"encoding/json"
	. "github.com/golangee/architecture/arc/adl"
)

const licenseExample = `Copyright 2021 Torben Schinke

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`

func createWorkspace() *Project {
	return NewProject("supportiety").
		PutGlossary("supportiety/tickets", "...describes the bounded context around anything in the error reporting context treated as a ticket.").
		AddModules(
			NewModule("server").
				SetLicense(licenseExample).
				SetGenerator(
					NewGenerator().
						SetOutDir("../../testdata/workspace/server").
						SetGo(NewGolang().SetModName("github.com/golangee/architecture/testdata/workspace/server")),
				).
				SetDomain(
					NewDomain("supportiety/tickets").
						AddCore(
							NewPackage("", "").AddRepositories(
								NewInterface("Tickets", "...provides CRUD access to Tickets.").
									AddMethods(
										NewMethod("CreateTicket", "...creates a Ticket.").
											AddIn("id", "...the unique ticket id", "uuid!").
											AddOut("", "...if anything goes wrong", "error!"),
									),
							),

							NewPackage("chat", "...is a supporting subdomain about ticket chats.").AddRepositories(
								NewInterface("Chats", "...provides CRUD access to Chats."),
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
