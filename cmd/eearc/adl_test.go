package eearc

import (
	"fmt"
	"github.com/golangee/architecture/arc"
	"github.com/golangee/src/render"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
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

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasSuffix(cwd, "eearc") {
		t.Fatalf("unexpected directory location: %v", cwd)
	}

	if err := render.Write(cwd, a); err != nil {
		t.Fatal(err)
	}

	// directly test our generated stuff
	cmd := exec.Command("go", "test", "-cover", "./...")
	cmd.Dir = filepath.Clean("../../testdata/workspace/server")
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	t.Log(string(out))
	if err != nil {
		t.Fatal(err)
	}
}
