package author

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	service AuthorService
}

type response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func NewAuthorHandler(service AuthorService) AuthorHandler {
	return AuthorHandler{service: service}
}

// GetAuthorByID godoc
// @Summary Get author
// @Description Get author by ID
// @Tags author
// @Produce  json
// @Param id path string true "Author ID"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Failure 404 {object} response
// @Router /authors/{id} [get]
func (h AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	var author Author
	idPar := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idPar)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response{Error: "Wrong type of ID"})
		return
	}
	author, err = h.service.GetAuthor(id)
	if author.ID == 0 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response{Error: "ID not found"})
		return
	}

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response{Error: "ID not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Data: author})

}
