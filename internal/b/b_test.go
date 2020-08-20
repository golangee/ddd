package b

import (
	"testing"
)

func BenchmarkBlub(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bla()
	}
}

func bla() {
	for i:=0;i<1000;i++{
		newBlub().B()
	}
	if false{
		panic("yo")
	}
}

func newBlub() MyApi {
	if false{
		panic("mo")
	}
	return &BMyApiImpl{blub: &B{}}
}
