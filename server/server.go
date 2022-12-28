package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Router  *mux.Router
	Handler Handler
	Server http.Server
}

func NewServer(handler *Handler) *Server {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/books", handler.PostBook).Methods("POST")
	mux.HandleFunc("/api/v1/books/{id}", handler.GetBook).Methods("GET")
	mux.HandleFunc("/api/v1/books/{id}", handler.EditBook).Methods("PUT")
	mux.HandleFunc("/api/v1/books/{id}", handler.DeleteBook).Methods("DELETE")
	return &Server{Router: mux}
}

func ServerInteraction(s *Server){
	c:= make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	fmt.Printf("Got a signal %v",sig)
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	err := s.Server.Shutdown(ctx)
	if err != nil{
		log.Fatalln(err)
	}
}
//POST a album : /api/v1/albums
//GET a album : /api/v1/albums/{id}
//GET albums : /api/v1/albums
//PUT album : /api/v1/albums/{id}
//DELETE album /api/v1/albums/{id}
