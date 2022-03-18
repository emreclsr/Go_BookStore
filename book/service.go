package book

import (
	"fmt"
	"os"
)

type service struct {
	repo repository
}

func NewService(repo repository) *service {
	return &service{repo: repo}
}

func (s service) List() {
	bookList := s.repo.GetAll()
	for _, i := range bookList {
		if i.Deleted == false {
			fmt.Println(i.ID, " ", i.Name)
		}
	}
}

func (s service) Search(name string) {
	list := SearchByWord(name, s.repo.Books)
	for _, i := range list {
		fmt.Println(i.Name)
	}
	if len(list) == 0 {
		fmt.Println("cannot found any record")
	}
}

func (s service) Getbook(id int) {
	book, index := SearchByID(id, s.repo.Books)
	if index == -1 {
		fmt.Println("cannot found any record with this id")
		os.Exit(0)
	} else {
		PrettyPrint(book)
	}
}

func (s service) Sell(id, quantity int, list []Book) {
	book, index := SearchByID(id, list)
	if index == -1 {
		fmt.Println("cannot found any record with this id")
		os.Exit(0)
	}

	if book.Stock < quantity {
		fmt.Printf("there is not enough stock for this book, only %v left", book.Stock)
		os.Exit(0)
	} else {
		book.Stock = book.Stock - quantity
		s.repo.Update(index, book)
		PrettyPrint(book)
	}
}

func (s service) Delete(id int) {
	book, index := SearchByID(id, s.repo.Books)
	if index == -1 {
		fmt.Println("cannot found any record with this id")
		os.Exit(0)
	} else {
		book.Deleted = true
		s.repo.Update(index, book)
		PrettyPrint(book)
	}
}
