package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmsaluno/my-notes/internal/db"
	handlers "github.com/lucasmsaluno/my-notes/internal/handler"
	"github.com/lucasmsaluno/my-notes/internal/repository"
	"github.com/lucasmsaluno/my-notes/internal/service"
	"github.com/rs/cors"
)

func main() {
	sqlite := db.InitDB("internal/db/notes.db")
	repo := repository.NewSQLiteNoteRepo(sqlite)
	service := service.NewNoteService(repo)
	handler := handlers.NewNoteHandler(service)

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	h := c.Handler(router)

	router.HandleFunc("/notes", handler.GetNotes).Methods("GET")
	router.HandleFunc("/notes", handler.CreateNote).Methods("POST")
	router.HandleFunc("/notes/{id:[0-9]+}", handler.UpdateNote).Methods("PUT")
	router.HandleFunc("/notes/{id:[0-9]+}", handler.DeleteNote).Methods("DELETE")

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}
