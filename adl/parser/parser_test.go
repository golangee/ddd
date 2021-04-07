package parser

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	mustParse := []string{
		`
module github.com/worldiety/supportiety {

	package what/an/evil/import.de/rust/::ticket {
	
		claim worldiety.com/supportiety/ManageTickets
		struct blub {
			msg other::stlib.com/evil/string
			upsilag bool
		}
	
		interface abc {
			doJob
		}

	}
}

requirements worldiety.com/supportiety

epic ManageTickets
"As a SupportietyAdmin or Application 
I want to manage tickets
so that I can submit or delete new incidents."

story AdminDeletesTicket
"As a SupportietyAdmin 
I want to delete tickets from a user identified by his SecId, 
so that I can comply to the DSGVO/GDPR."

story AdminCreatesTicket
"As a SupportietyAdmin 
I want to create tickets in case a user cannot transmit his data automatically."

epic AdminDeletesTicket 
"As a SupportietyAdmin 
I want to delete tickets from a user identified by his SecId, 
so that I can comply to the DSGVO/GDPR"


`,
	}

	for _, s := range mustParse {
		ast, err := Parse(s)
		fmt.Println(toString(ast))

		if err != nil {
			t.Fatal(err)
		}
	}
}

func toString(i interface{}) string {
	buf, err := json.MarshalIndent(i, " ", " ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
