package book

import (
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type BookRepository interface {
	Get(id int) (*Book, error)
	GetAll() ([]Book, error)
	Create(author Book) (int, error)
	Update(id int, book Book) error
	Delete(id int) error
}

// Compile time proof of interface implementation
var _ BookRepository = repository{}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo repository) Get(id int) (*Book, error) {
	book := &Book{ID: id}
	err := repo.db.First(book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

func (repo repository) GetAll() ([]Book, error) {
	var books []Book
	err := repo.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (repo repository) Create(book Book) (int, error) {
	err := repo.db.Create(&book).Error
	if err != nil {
		return 0, err
	}
	return int(book.ID), nil
}

func (repo repository) Update(id int, newbook Book) error {
	book := &Book{ID: id}
	err := repo.db.First(&book).Error
	if err != nil {
		return err
	}
	book.ID = newbook.ID
	book.Name = newbook.Name
	book.ISBN = newbook.ISBN
	book.Deleted = newbook.Deleted
	book.PageNumber = newbook.PageNumber
	book.Price = newbook.Price
	book.Stock = newbook.Stock
	book.StockCode = newbook.StockCode

	err = repo.db.Save(&book).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (repo repository) Delete(id int) error {
	book := &Book{ID: id}
	err := repo.db.Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
