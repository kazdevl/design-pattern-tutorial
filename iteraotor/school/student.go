package school

import (
	"fmt"

	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_item"
)

type Student struct {
	name   string
	number int
}

var _ if_item.IFItem = (*Student)(nil)

func NewStudent(name string, number int) *Student {
	return &Student{name: name, number: number}
}

func (s *Student) Name() string {
	return fmt.Sprintf("%d番号:%s", s.number, s.name)
}
