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
	bookList, err := s.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range bookList {
		if i.Deleted == false {
			fmt.Println(i.ID, " ", i.Name)
		}
	}
}

func (s service) Search(name string) {
	books, err := s.repo.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	list := SearchByWord(name, books)
	for _, i := range list {
		fmt.Println(i.Name)
	}
	if len(list) == 0 {
		fmt.Println("cannot found any record")
	}
}

func (s service) GetBook(id int) {
	book, err := s.repo.Get(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	PrettyPrint(*book)
}

func (s service) GetAuthor(id int) string {
	var authorName string
	s.repo.db.Raw("SELECT AUTHORS.NAME FROM AUTHORS WHERE id = ?", id).Scan(&authorName)
	return authorName
}

func (s service) Sell(id, quantity int) {
	book, err := s.repo.Get(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if book.Stock < quantity {
		fmt.Printf("there is not enough stock for this book, only %v left", book.Stock)
		os.Exit(0)
	} else {
		book.Stock = book.Stock - quantity
		s.repo.Update(book.ID, *book)
		PrettyPrint(*book)
	}
}

func (s service) Delete(id int) {
	book, err := s.repo.Get(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	book.Deleted = true
	s.repo.Update(book.ID, *book)

	PrettyPrint(*book)
}

func (s service) GetByAuthor(author string) string {
	var getauthor Book
	rows, err := s.repo.db.Raw("SELECT B.NAME, A.NAME FROM AUTHORS A JOIN BOOKS B on A.ID=b.AUTHOR_ID WHERE A.NAME LIKE '%" + author + "%'").Rows()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for rows.Next() {
		err := rows.Scan(&getauthor.Name, &getauthor.Author.Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(getauthor.Name, " by ", getauthor.Author.Name)
	}
	return ""
}

func (s service) GetByBook(book string) string {
	var getBook Book
	rows, err := s.repo.db.Raw("SELECT B.NAME, A.NAME FROM AUTHORS A JOIN BOOKS B on A.ID=b.AUTHOR_ID WHERE B.NAME LIKE '%" + book + "%'").Rows()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&getBook.Name, &getBook.Author.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(getBook.Name, " by ", getBook.Author.Name)
	}
	return ""
}

func (s service) PriceFilter(min, max int) string {
	var bookfilter Book
	rows, err := s.repo.db.Raw("SELECT B.NAME, A.NAME, B.PRICE FROM AUTHORS A JOIN BOOKS B on A.ID=B.AUTHOR_ID WHERE B.PRICE BETWEEN ? AND ? ORDER BY B.PRICE DESC", min, max).Rows()
	fmt.Println("Books with price between", min, "and", max)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&bookfilter.Name, &bookfilter.Author.Name, &bookfilter.Price)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bookfilter.Name, " by ", bookfilter.Author.Name, "price:", bookfilter.Price)
	}

	return ""
}
