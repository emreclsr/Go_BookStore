package main

import (
	"github.com/emreclsr/book/author"
	"github.com/emreclsr/book/book"
	"github.com/emreclsr/book/db"
	_ "github.com/emreclsr/book/docs"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Book API
// @version 1.0
// @description This is a sample an exercise written in Go.

// @contact.name Emre ÇALIŞIR

// @host localhost:6363

// @accept json
// @produce json

// @BasePath /api
// @schemes http https

func main() {

	// Set env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//// Connect to db
	//db, err := db.Connect()
	//if err != nil {
	//	log.Fatal("Postgres cannot init: ", err)
	//}

	repo, err := db.NewRepositories()
	if err != nil {
		log.Fatal("Postgres cannot init: ", err)
	}
	//handler := book.NewHandler(service)

	bookService := book.NewBookService(repo.Book)
	authorService := author.NewAuthorService(repo.Author)

	bookHandler := book.NewBookHandler(bookService)
	authorHandler := author.NewAuthorHandler(authorService)

	router := mux.NewRouter()

	router.HandleFunc("/api/books", bookHandler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/search/{word}", bookHandler.SearchBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", bookHandler.GetBookByID).Methods(http.MethodGet)
	router.HandleFunc("/api/authors/{id}", authorHandler.GetAuthorByID).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}/sell/{quantity}", bookHandler.SellBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":6363", router)

	//reading csv file after that turn readed data into structs and insert the struct via gorm Automigrate tool.
	//repo := book.NewRepository(db)
	//repo.Migration()
	//
	//test := utils.CsvToStruct("./dummy.csv")
	//for _, i := range test {
	//	id, err := repo.Create(i)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(id)
	//}
}
