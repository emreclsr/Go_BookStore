package main

import (
	"github.com/emreclsr/book/book"
	"github.com/emreclsr/book/utils"
)

func main() {
	book1 := &book.Book{ID: 1, Name: "A Tale of Two Cities", PageNumber: 416, Stock: 5,
		Price: 7, StockCode: 5, ISBN: "978-0451530578", Author: book.Author{ID: 1, Name: "Charles Dickens"}, Deleted: false}
	book2 := &book.Book{ID: 2, Name: "Harry Potter and the Philosopher's Stone", PageNumber: 272, Stock: 5,
		Price: 20, StockCode: 5, ISBN: "978-1408866191", Author: book.Author{ID: 2, Name: "J.K. Rowling"}, Deleted: false}
	book3 := &book.Book{ID: 3, Name: "The Hobbit", PageNumber: 320, Stock: 5,
		Price: 8, StockCode: 5, ISBN: "978-0618260300", Author: book.Author{ID: 3, Name: "J.R.R. Tolkien"}, Deleted: false}
	book4 := &book.Book{ID: 4, Name: "The Lord of the Rings", PageNumber: 1536, Stock: 5,
		Price: 22, StockCode: 5, ISBN: "978-0358439196", Author: book.Author{ID: 3, Name: "J.R.R. Tolkien"}, Deleted: false}
	var list []book.Book
	list = append(list, *book1, *book2, *book3, *book4)
	utils.Serve(&list)

}
