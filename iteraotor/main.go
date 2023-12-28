package main

import (
	"github.com/kazdev/design-pattern-tutorial/iteraotor/book"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/if_iterator"
	"github.com/kazdev/design-pattern-tutorial/iteraotor/school"
)

func main() {
	var (
		iterator if_iterator.IFIterator
	)

	bookshelf := book.NewBookShelf()
	bookshelf.AppendBook(book.NewBook("Around the World in 80 Days"))
	bookshelf.AppendBook(book.NewBook("Bible"))
	bookshelf.AppendBook(book.NewBook("Cinderella"))
	bookshelf.AppendBook(book.NewBook("Daddy-Long-Legs"))

	iterator = bookshelf.Iterator()
	for iterator.HasNext() {
		book := iterator.Next()
		println(book.Name())
	}

	class := school.NewClass()
	class.AddStudent(school.NewStudent("Alice", 210))
	class.AddStudent(school.NewStudent("Bob", 100))
	class.AddStudent(school.NewStudent("Charlie", 130))

	iterator = class.Iterator()
	for iterator.HasNext() {
		student := iterator.Next()
		println(student.Name())
	}
}
