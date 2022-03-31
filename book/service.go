package book

import (
	"fmt"
)

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return bookService{repo: repo}
}

type BookService interface {
	GetBook(id int) (Book, error)
	GetBooks() ([]Book, error)
	SearchBook(name string) ([]Book, error)
	SellBook(id, quantity int) error
	DeleteBook(id int) error
}

// Compile time proof of interface implementation
var _ BookService = bookService{}

func (s bookService) GetBook(id int) (Book, error) {
	book, err := s.repo.Get(id)
	if err != nil {
		return Book{}, err
	}
	return *book, nil
}

func (s bookService) GetBooks() ([]Book, error) {
	bookList, err := s.repo.GetAll()
	var result []Book
	for _, book := range bookList {
		if book.Deleted == false {
			result = append(result, book)
		}
	}
	if err != nil {
		return []Book{}, err
	}
	return result, nil
}

func (s bookService) SearchBook(name string) ([]Book, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return []Book{}, err
	}
	list := SearchByWord(name, books)
	return list, nil
}

func (s bookService) SellBook(id, quantity int) error {
	book, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	if book.Stock < quantity {
		return fmt.Errorf("Not enough stock")
	} else {
		book.Stock = book.Stock - quantity
		s.repo.Update(book.ID, *book)
		return nil
	}
}

func (s bookService) DeleteBook(id int) error {
	book, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	book.Deleted = true
	err = s.repo.Update(book.ID, *book)
	if err != nil {
		return err
	}
	return nil
}
