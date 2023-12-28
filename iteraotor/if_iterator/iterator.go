package if_iterator

import "github.com/kazdev/design-pattern-tutorial/iteraotor/if_item"

type IFIterator interface {
	HasNext() bool
	Next() if_item.IFItem
}
