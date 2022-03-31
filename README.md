## Go Book Store

This program provides functions for events that may occur in a book store.
Functions can be running with flags via CLI.
If the mode flag is empty or wrong, the program return usage.

*This program is an example of an exercise written in Go language.*

### DB Schema

![DB Schema](https://cdn.discordapp.com/attachments/519918508998656028/959028680330469386/unknown.png)

 - **book list**
```
localhost:6363/api/books
```
This endpoint return all the books and their informations.

 - **delete (soft)**
```
localhost:6363/api/books/{id}
```
This endpoint change delete properties of book which id is given. Command is not actually delete book. Only changes one parameter (Delete) and book becomes invisible.

 - **get book by id**
```
localhost:6363/api/books/{id}
```
This endpoint get book by given id.

 - **get books by name**
```
localhost:6363/api/books/search/{word}
```
This endpoint get books by given name.

 - **sell**
```
localhost:6363/api/books/{id}/sell/{quantity}
```
This endpoint sell the given quantity for the book which id is given. Command also compare quantity and book stock.

- **get author by id**
```
localhost:6363/api/authors/{id}
```
This endpoint get author by given id.
