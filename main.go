package main

import (
	"github.com/emreclsr/book/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Set env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	////reading csv file after that turn readed data into structs and insert the struct via gorm Automigrate tool.
	//
	//db, err := db.Connect()
	//if err != nil {
	//	log.Fatal("Postgres cannot init: ", err)
	//}
	//log.Print("Postgres connected")
	//
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
	utils.Serve()
}
