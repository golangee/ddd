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
		... "shows how to declare a struct data type"	
		struct structname {
				
			... "is another message"	
			msg other::stlib.com/evil/string
			
			... "is a flag"
			upsiflag bool
		}
	
		... "is an abc interface"
		interface abc {
			... "does the job."
			doJob(x string, y string)
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

scenario DoubleDelete
"Given I'm a SupportietyAdmin
when I delete the same ticket twice
then I want a message telling me that its not possible."

epic AnotherEpic 
"blablabla"

story EpicStory1
"blablub"

story EpicStory2
"## Should markdown be allowed?
What about 
 * pointless
 * points

And tables

| value | size |
|-------|------|
| c     | d    |
|       |      |
|       |      |

"
scenario OfEpic2-Scenario1
"scenario 1 of epic 2 "

scenario OfEpic2-Scenario2
"scenario 2 of epic 2 "

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
