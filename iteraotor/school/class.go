package school

import (
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_aggregte"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"
)

type Class struct {
	Students []*Student
}

var _ if_aggregte.IFAggregate = (*Class)(nil)

func NewClass() *Class {
	return &Class{}
}

func (c *Class) GetStduent(index int) *Student {
	return c.Students[index]
}

func (c *Class) AddStudent(student *Student) {
	c.Students = append(c.Students, student)
}

func (c *Class) StudentCount() int {
	return len(c.Students)
}

func (c *Class) Iterator() if_iterator.IFIterator {
	return NewClassIterator(c)
}
