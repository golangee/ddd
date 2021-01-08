package aclyaml

import (
	"fmt"
	"testing"
)

const test = `
- Book:
    userstories:
      - 001
      - 002
      - 042

    fields2: {
      id: string
    }

    fields:
      - ID: string! # stdlib type
      - Authors: slice[shared/User] # list or array or slice of shared/User type
      - OtherUser: ?User # pointer which may be nil
      - AndEvenMore: !User # pointer which may NOT be nil

- User:
    - SpecialID: int32
    - Name: string`

func TestDecode(t *testing.T) {
	pkg,err := Parse("/Users/tschinke/git/github.com/golangee/architecture.git/ddd/v2/example")
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	fmt.Println(pkg)
}
