package book

import "github.com/kazdev/design-pattern-tutorial/iteraotor/if_item"

type Book struct {
	name string
}

var _ if_item.IFItem = (*Book)(nil)

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b *Book) Name() string {
	return b.name
}
