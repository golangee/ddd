package arch

import (
	"errors"
	"github.com/golangee/architecture/objn"
	"testing"
)

func TestBuild(t *testing.T) {
	err := Build("/Users/tschinke/git/github.com/golangee/architecture.git/ddd/v2/example")
	if err != nil {
		var locErr objn.PosError
		if errors.As(err, &locErr) {
			t.Fatal(locErr.Explain())
		}

		t.Fatal(err)
	}

}
