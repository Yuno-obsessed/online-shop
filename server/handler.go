package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	book "practice/model"
	"practice/mysql"
	"strconv"
)

type Handler struct {
	Conn *mysql.DbConn
}

func NewHandler(conn *mysql.DbConn) *Handler {
	return &Handler{Conn: conn}
}

func (h *Handler) PostBook(rw http.ResponseWriter, r *http.Request) {
	var newBook book.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	_, err = h.Conn.Insert(newBook)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Server error", http.StatusInternalServerError)
		return
	}
	//json.NewEncoder(rw).Encode(newBook)

	rw.Write([]byte("Book was successfully added"))
}

func (h *Handler) GetBook(rw http.ResponseWriter, r *http.Request) {

	requestId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(requestId)
	target, err := h.Conn.Get(id)
	if err != nil {
		http.Error(rw, "Wrong book id", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(target)
	if err != nil {
		http.Error(rw, "Server Error", http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("Book with id " + requestId))
}

func (h *Handler) EditBook(rw http.ResponseWriter, r *http.Request) {

	var newBook book.Book
	requestId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(requestId)
	if err != nil {
		http.Error(rw, "Wrong book id", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(rw, "Wrong post request data", http.StatusBadRequest)
		return
	}
	_, err = h.Conn.Edit(id, newBook)
	if err != nil {
		http.Error(rw, "Wrong book id", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("Book with uuid " + requestId + " was successfully edited/"))
}

func (h *Handler) DeleteBook(rw http.ResponseWriter, r *http.Request) {
	requestedId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(requestedId)
	if err != nil {
		http.Error(rw, "Wrong book id", http.StatusBadRequest)
		return
	}
	_, err = h.Conn.Delete(id)
	if err != nil {
		http.Error(rw, "Wrong book id", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode("book with uuid" + requestedId + " was deleted.")
	rw.WriteHeader(200)
}
