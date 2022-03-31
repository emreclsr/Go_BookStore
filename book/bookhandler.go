package book

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	service BookService
}

type response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func NewBookHandler(service BookService) BookHandler {
	return BookHandler{service: service}
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags book
// @Produce json
// @Success 200 {object} response
// @Failure 500 {object} response
// @Router /books [get]
func (h BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := h.service.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: bookList})
}

// SearchBooks godoc
// @Summary Search book
// @Description Search book by given words
// @Tags book
// @Produce  json
// @Param word path string true "Search word"
// @Success 200 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Router /books/search/{word} [get]
func (h BookHandler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	searchWord := mux.Vars(r)["word"]
	word := string(searchWord)

	bookList, err := h.service.SearchBook(word)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Error: err.Error()})
		return
	}
	if len(bookList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{Error: "No books found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: bookList})
}

// GetBookByID godoc
// @Summary Get book
// @Description Get book by ID
// @Tags book
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Failure 404 {object} response
// @Router /books/{id} [get]
func (h BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	var book Book
	idPar := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idPar)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response{Error: "Wrong type of ID"})
		return
	}
	book, err = h.service.GetBook(id)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response{Error: "ID not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: book})
}

// SellBook godoc
// @Summary Sell book
// @Description Sell book by ID and quantity
// @Tags book
// @Produce  json
// @Param id path string true "Book ID"
// @Param quantity path string true "Quantity"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /books/{id}/sell/{quantity} [put]
func (h BookHandler) SellBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response{Error: "Wrong type of ID"})
		return
	}
	givenQuantity := mux.Vars(r)["quantity"]
	quantity, err := strconv.Atoi(givenQuantity)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response{Error: "Wrong type of quantity"})
		return
	}
	err = h.service.SellBook(id, quantity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: "Book sold"})
}

// DeleteBook godoc
// @Summary Delete book
// @Description Delete book by ID
// @Tags book
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /books/{id} [delete]
func (h BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response{Error: "Wrong type of ID"})
		return
	}
	err = h.service.DeleteBook(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: "Book deleted"})
}
