package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucasmsaluno/my-notes/internal/service"
)

type NoteHandler struct {
	service *service.NoteService
}

func NewNoteHandler(service *service.NoteService) *NoteHandler {
	return &NoteHandler{service}
}

func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetAllNotes()
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	note, err := h.service.CreateNote(input.Content)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var input struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateNote(id, input.Content); err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteNote(id); err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
