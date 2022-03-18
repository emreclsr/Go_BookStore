package book

import (
	"errors"
)

type repository struct {
	Books []Book //Db connection will be added here
}

func NewRepository(books []Book) *repository {
	return &repository{Books: books}
}

func (repo repository) Create(book Book) error {
	repo.Books = append(repo.Books, book)
	return nil
}

func (repo repository) GetAll() []Book {
	var list []Book
	for _, i := range repo.Books {
		if i.Deleted == false {
			list = append(list, i)
		}
	}
	return list
}

func (repo repository) Get(id int) (Book, error) {

	newbook, index := SearchByID(id, repo.Books)
	if index != -1 {
		return newbook, nil
	}
	return Book{}, errors.New("there is no book with given id")
}

func (repo repository) Update(id int, newbook Book) error {
	_, index := SearchByID(id, repo.Books)
	if index == -1 {
		return errors.New("error update")
	}
	repo.Books[index] = newbook
	return nil
}

func (repo repository) Delete(id int) error {
	_, index := SearchByID(id, repo.Books)
	if index != -1 {
		repo.Books[index].Deleted = true
		return nil
	}
	return errors.New("there is no element with given id")
}
