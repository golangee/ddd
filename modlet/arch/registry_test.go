package arch

import "testing"

func TestBuild(t *testing.T) {
	err := Build("/Users/tschinke/git/github.com/golangee/architecture.git/ddd/v2/example")
	if err != nil {
		t.Fatal(err)
	}

}
