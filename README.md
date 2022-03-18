## Homework | Week 2
### Book Lister and Searcher

The code lists the book names in the file and queries whether the searched book name is in the list.

*Script is an example of an exercise written in Go language.*

 - **list command**
```
go run main.go list
```
This command list all the books in the code file.

 - **search command**
```
go run main.go <bookName>
```
This command searches the given book information and returns whether it exists or not. This command is not case sensetive.

 - **help command**
```
go run main.go help
```
The help command shows the explanation needed to use the application. It also works in error situations.
