## Go Book Store

This program provides functions for events that may occur in a book store.
Functions can be running with flags via CLI.
If the mode flag is empty or wrong, the program return usage.

*This program is an example of an exercise written in Go language.*

### DB Schema

![DB Schema](https://cdn.discordapp.com/attachments/561819205742362627/957000910666616863/unknown.png)

 - **list command**
```
go run main.go --mode=list
```
This command list all the books and id informations in the code file.

 - **delete (soft) command**
```
go run main.go --mode=delete --key1=id
```
This command change delete properties of book which id is given. Command is not actually delete book. Only changes one parameter (Delete) and book becomes invisible.

 - **getByName command**
```
go run main.go --mode=getByName --key1=author or book --key2=words
```
This command get books and authors by given name.

 - **getById command**
```
go run main.go --mode=getById --key1=author or book --key2=id
```
This command get books and authors by given id.

 - **buy command**
```
go run main.go --mode=buy --key1=id --key2=quantity
```
This command buy the given quantity for the book which id is given. Command also compare quantity and book stock.

- **filter command**
```
go run main.go --mode=filter --key1=min_price --key2=max_price
```
This command filter books by given price range.
