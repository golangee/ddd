package golang

import (
	"strconv"
	"testing"
)

func Test_sqlChecksum(t *testing.T) {
	tests := []struct {
		statements []string
		want       string
	}{
		{statements: []string{"abc"}, want: "3a985da74fe225b2045c172d6bd390bd"},
		{statements: []string{"abc "}, want: "3a985da74fe225b2045c172d6bd390bd"},
		{statements: []string{"abc \n"}, want: "3a985da74fe225b2045c172d6bd390bd"},

		{statements: []string{" a b c"}, want: "fd3f2cb4f0d2dff4bc0dd177277c32c7"},
		{statements: []string{" a\nb\nc "}, want: "fd3f2cb4f0d2dff4bc0dd177277c32c7"},
		{statements: []string{"a  \nb    	\nc"}, want: "fd3f2cb4f0d2dff4bc0dd177277c32c7"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sqlChecksum(tt.statements); got != tt.want {
				t.Errorf("sqlChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
