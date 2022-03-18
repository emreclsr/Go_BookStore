package utils

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/emreclsr/book/book"
)

type Input struct {
	mode     string
	data     string
	quantity string
}

func Usage() {
	fmt.Println(`
Welcome
There are 5 option for Bookstore
For listing all the records you can use go run main.go --mode=list
For searching object go run main.go --mode=search --entry=words
For deleting object go run main.go --mode=delete --entry=id
For looking for spesific record go run main.go --mode=getbyid --entry=id
For buy object go run main.go --mode=buy --entry=id --quantity=quantity`)
}

func GetInput() Input {
	mode := flag.String("mode", "help", "functions -list,-delete,-search,-getbyid,-buy")
	data := flag.String("entry", "default", "name for list otherwise id")
	quantity := flag.String("quantity", "default", "quantity for buy")
	flag.Parse()
	input := Input{mode: *mode, data: *data, quantity: *quantity}
	return input
}

func Serve(books *[]book.Book) {
	repo := book.NewRepository(*books)
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
	case "search":
		{
			service.Search(inputs.data)
		}
	case "getbyid":
		{
			id, err := strconv.Atoi(inputs.data)
			if err != nil {
				fmt.Println("please check entered id")
				os.Exit(0)
			}
			service.Getbook(id)
		}
	case "buy":
		{
			id, err1 := strconv.Atoi(inputs.data)
			quantity, err2 := strconv.Atoi(inputs.quantity)
			if err1 != nil || err2 != nil {
				fmt.Println("please check entered id or quantity")
				os.Exit(0)
			} else if quantity <= 0 {
				fmt.Println("please check entered quantity")
			} else {
				service.Sell(id, quantity, *books)
			}
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
