package school

import (
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_item"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"
)

type ClassIterator struct {
	class    *Class
	position int
}

var _ if_iterator.IFIterator = (*ClassIterator)(nil)

func NewClassIterator(class *Class) *ClassIterator {
	return &ClassIterator{
		class:    class,
		position: 0,
	}
}

func (ci *ClassIterator) HasNext() bool {
	return ci.position < ci.class.StudentCount()
}

func (ci *ClassIterator) Next() if_item.IFItem {
	student := ci.class.GetStduent(ci.position)
	ci.position++
	return student
}
