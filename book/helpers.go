package book

import (
	"fmt"
	"strings"
)

func SearchByID(id int, list []Book) (Book, int) {
	for k, i := range list {
		if i.ID == id {
			return i, k
		}
	}
	return Book{}, -1
}

func SearchByWord(word string, list []Book) []Book {
	var NewBooks []Book
	for _, i := range list {
		if strings.Contains(strings.ToLower(i.Name), strings.ToLower(word)) {
			NewBooks = append(NewBooks, i)
		}
	}
	return NewBooks
}

func PrettyPrint(book Book) {
	fmt.Printf(`
Book ID: %v
Book Name: %v
Book PageNumber: %v
Book Stock: %v
Book Price: %v
Book StockCode: %v
Book ISBN: %v
Author ID: %v
Deleted: %v`, book.ID, book.Name, book.PageNumber, book.Stock, book.Price, book.StockCode, book.ISBN, book.AuthorID, book.Deleted)
}
