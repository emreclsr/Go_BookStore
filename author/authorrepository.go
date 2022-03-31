package author

import (
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type AuthorRepository interface {
	Get(id int) (Author, error)
	GetAll() ([]Author, error)
	Create(author Author) (int, error)
	Update(id int, author Author) error
	Delete(id int) error
}

// Compile time proof of interface implementation
var _ AuthorRepository = repository{}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return repository{db: db}
}

func (repo repository) Get(id int) (Author, error) {
	author := Author{ID: id}
	err := repo.db.First(&author).Error
	if err != nil {
		return Author{}, err
	}
	return author, nil
}

func (repo repository) GetAll() ([]Author, error) {
	var authors []Author
	err := repo.db.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (repo repository) Create(author Author) (int, error) {
	err := repo.db.Create(&author).Error
	if err != nil {
		return 0, err
	}
	return int(author.ID), nil
}

func (repo repository) Update(id int, newauthor Author) error {
	author := &Author{ID: id}
	err := repo.db.First(&author).Error
	if err != nil {
		return err
	}
	author.ID = newauthor.ID
	author.Name = newauthor.Name

	err = repo.db.Save(&author).Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (repo repository) Delete(id int) error {
	author := &Author{ID: id}
	err := repo.db.Delete(&author).Error
	if err != nil {
		return err
	}
	return nil
}
