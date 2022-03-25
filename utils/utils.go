package utils

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/emreclsr/book/book"
	"github.com/emreclsr/book/db"
)

type Input struct {
	mode  string
	data  string
	value string
}

func Usage() {
	fmt.Println(`
Welcome
There are 6 option for Bookstore
For listing all the records you can use go run main.go --mode=list
For deleting object                     go run main.go --mode=delete    --key1=id
For looking record                      go run main.go --mode=getByName --key1=author or book --key2=words
For searching book or author with id    go run main.go --mode=getById   --key1=author or book --key2=id
For buy object                          go run main.go --mode=buy       --key1=id             --key2=quantity
For filtering depend on their price     go run main.go --mode=filter    --key1=min_price      --key2=max_price`)
}

func GetInput() Input {
	mode := flag.String("mode", "help", "functions -list,-delete,-getByName,-getById,-buy,-filter")
	data := flag.String("key1", "default", "for usage --help")
	value := flag.String("key2", "default", "for usage --help")
	flag.Parse()
	input := Input{mode: *mode, data: *data, value: *value}
	return input
}

func Serve() {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	repo := book.NewRepository(db)
	service := book.NewService(*repo)
	inputs := GetInput()

	switch inputs.mode {
	case "list":
		{
			service.List()
		}
	case "delete":
		{
			id, err := strconv.Atoi(inputs.data)
			if err != nil {
				fmt.Println("please check entered id")
				os.Exit(0)
			}
			service.Delete(id)
		}
	case "getById":
		{
			entry := inputs.data
			id, err := strconv.Atoi(inputs.value)
			if err != nil {
				fmt.Println("please check entered id")
				os.Exit(0)
			}
			if entry == "author" {
				fmt.Println(service.GetAuthor(id))
			} else if entry == "book" {
				service.GetBook(id)
			} else {
				fmt.Println("please check entered entry")
				os.Exit(0)
			}

		}
	case "getByName":
		{
			entry := inputs.data
			value := inputs.value
			if entry == "author" {
				service.GetByAuthor(value)
			} else if entry == "book" {
				service.GetByBook(value)
			} else {
				fmt.Println("please check entered entry")
				os.Exit(0)
			}

		}
	case "buy":
		{
			id, err := strconv.Atoi(inputs.data)
			if err != nil {
				fmt.Println("please check entered id")
				os.Exit(0)
			}
			quantity, err := strconv.Atoi(inputs.value)
			if err != nil {
				fmt.Println("please check entered quantity")
				os.Exit(0)
			} else if quantity <= 0 {
				fmt.Println("please check entered quantity")
			} else {
				service.Sell(id, quantity)
			}
		}
	case "filter":
		{
			min, err := strconv.Atoi(inputs.data)
			if err != nil {
				fmt.Println("please check entered min price")
				os.Exit(0)
			}
			max, err := strconv.Atoi(inputs.value)
			if err != nil {
				fmt.Println("please check entered max price")
				os.Exit(0)
			}
			service.PriceFilter(min, max)
		}

	case "help":
		{
			Usage()
		}
	default:
		{
			Usage()
		}

	}

}

//service.Search(inputs.data)

func CsvToStruct(path string) []book.Book {
	var newbook book.Book
	var books []book.Book
	// Creating a connection via csv files
	csvFile, err := os.Open(path)
	defer csvFile.Close()

	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while csv read operation ")
	}
	//creating a *csv.reader
	reader := csv.NewReader(csvFile)
	//reading csv files via reader and method of reader
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		fmt.Println("something went wrong while reading csv file ")
	}
	//in csv file, each comma represent a column, with this knowledge every row has 12 column and each column parse where they are belong.
	for i, each := range csvData {
		if i > 0 {

			bookid, _ := strconv.Atoi(each[0])
			newbook.ID = bookid
			newbook.Name = each[1]
			newbook.PageNumber, _ = strconv.Atoi(each[2])
			newbook.Stock, _ = strconv.Atoi(each[3])
			newbook.Price, _ = strconv.ParseFloat(each[4], 64)
			newbook.StockCode, _ = strconv.Atoi(each[5])
			newbook.ISBN = each[6]
			newbook.Author.ID, _ = strconv.Atoi(each[7])
			newbook.Author.Name = each[8]
			newbook.Deleted, _ = strconv.ParseBool(each[9])
			books = append(books, newbook)
		}

	}

	return books
}
