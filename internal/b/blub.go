package b

import (
	"github.com/golangee/ddd/internal/a"
)

type BMyApiImpl struct {
	blub *B
}

func (m *BMyApiImpl) B() B {
	return *m.blub
}

func (m *BMyApiImpl) A() a.A {
	panic("implement me")
}
