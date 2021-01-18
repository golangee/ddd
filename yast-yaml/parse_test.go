package yastyaml

import (
	"fmt"
	"github.com/golangee/architecture/yast"
	yastjson "github.com/golangee/architecture/yast-json"
	"testing"
)

func TestParseDir(t *testing.T) {
	pkg, err := ParseDir("/Users/tschinke/git/github.com/golangee/architecture.git/ddd/v2/example")
	if err != nil {
		fmt.Println(yast.Explain(err))
		t.Fatal(err)
	}

	fmt.Println(yastjson.Sprint(pkg))
}
