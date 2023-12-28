package if_aggregte

import "github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"

type IFAggregate interface {
	Iterator() if_iterator.IFIterator
}
