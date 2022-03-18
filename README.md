## Go Book Store

This program provides functions for events that may occur in a book store.
Functions can be running with flags via CLI.
If the mode flag is empty or wrong, the program return usage.

*This program is an example of an exercise written in Go language.*

 - **list command**
```
go run main.go --mode=list
```
This command list all the books and id informations in the code file.

 - **search command**
```
go run main.go --mode=search --entry=words
```
This command searches the given words and return book information if book name contains given words. This command is not case sensetive.

 - **delete command**
```
go run main.go --mode=delete --entry=id
```
This command change delete properties of book which id is given. Command is not actually delete book. Only changes one parameter (Delete) and book becomes invisible.

 - **getbyid command**
```
go run main.go --mode=getbyid --entry=id
```
This command show book informations with the given id.

 - **buy command**
```
go run main.go --mode=buy --entry=id --quantity=quantity
```
This command buy the given quantity for the book which id is given. Command also compare quantity and book stock.
