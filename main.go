package main

import (
	"fmt"
	"os"
	"strings"
)

type Books struct {
	Name   string
	Author string
}

var bookList = []Books{
	{
		Name:   "A Tale of Two Cities",
		Author: "Charles Dickens",
	}, {
		Name:   "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
	}, {
		Name:   "The Da Vinci Code",
		Author: "Dan Brown",
	}, {
		Name:   "The Hobbit",
		Author: "J.R.R. Tolkien",
	}, {
		Name:   "The Little Prince",
		Author: "Antoine de Saint-Exupéry",
	}, {
		Name:   "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
	},
}

func (b Books) GetName() string {
	return b.Name
}

func Search(bookName string) {
	boolean := false
	bookNameLower := strings.ToLower(bookName)
	var libBook string

	for _, book := range bookList {
		if strings.ToLower(book.GetName()) == bookNameLower {
			boolean = true
			libBook = book.Name
			break
		}
		libBook = bookName
	}

	if boolean {
		fmt.Printf("The book is in the list: %s\n", libBook)
	} else {
		fmt.Printf("The book is not in the list: %s\n", bookName)
	}
}

func List(bList []Books) {
	for _, book := range bookList {
		fmt.Printf("%s\n", book.GetName())
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("go run main.go list")
	fmt.Println("go run main.go search <\"book name\">")
}

func main() {

	if len(os.Args) > 1 {

		if os.Args[1] == "list" {
			List(bookList)

		} else if os.Args[1] == "search" {

			if len(os.Args) < 3 {
				fmt.Println("Please enter a book name")
				help()

			} else {
				bookName := strings.Join(os.Args[2:], " ")
				Search(bookName)
			}
		} else if len(os.Args) > 1 && os.Args[1] == "help" {
			help()

		} else {
			fmt.Println("Please enter a valid command")
			help()
		}
	} else {
		fmt.Println("Please enter a valid command")
		help()
	}
}
