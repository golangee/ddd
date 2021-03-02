package modlet

import (
	"fmt"
	"github.com/golangee/architecture/yast"
	yastjson "github.com/golangee/architecture/yast-json"
	"testing"
)

func TestParseWorkspace(t *testing.T) {
	ws, err := ParseWorkspace("/Users/tschinke/git/github.com/golangee/architecture.git/ddd/v2/Supportiety")
	if err != nil {
		fmt.Println(yast.Explain(err))
		t.Fatal(err)
	}

	fmt.Println(yastjson.Sprint(ws.ADL))

	fmt.Printf("\n=========CODE========\n")
	artifact, err := ws.Render()
	if err != nil {
		fmt.Println(artifact)
		t.Fatal(err)
	}

	fmt.Println(artifact)
}
