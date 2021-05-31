package eearc

import (
	"fmt"
	"github.com/golangee/architecture/arc"
	"reflect"
	"testing"
)

func TestADLMarshalling(t *testing.T) {
	ws := createWorkspace()
	buf := toJson(ws)
	fmt.Println(buf)
	prj2 := wsFromJson(buf)
	if !reflect.DeepEqual(ws, prj2) {
		t.Fatalf("expected workspaces to be equal: %s", buf)
	}
}

func TestDummyProject(t *testing.T) {
	ws := createWorkspace()
	a, err := arc.Render(ws)
	fmt.Println(a)

	if err != nil {
		t.Fatal(err)
	}
}
