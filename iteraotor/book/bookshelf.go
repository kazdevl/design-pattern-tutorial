package book

import (
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_aggregte"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"
)

type BookShelf struct {
	books []*Book
}

var _ if_aggregte.IFAggregate = (*BookShelf)(nil)

func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0),
	}
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelf) Length() int {
	return len(bs.books)
}

func (bs *BookShelf) Iterator() if_iterator.IFIterator {
	return NewBookShelfIterator(bs)
}
