package book

import (
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_item"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"
)

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

var _ if_iterator.IFIterator = (*BookShelfIterator)(nil)

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.Length()
}

func (bsi *BookShelfIterator) Next() if_item.IFItem {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
