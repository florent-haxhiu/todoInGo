package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	db "florent-haxhiu/todoInGo/internal/database"
	"florent-haxhiu/todoInGo/internal/model"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	stringId := chi.URLParam(r, "noteId")

	noteId := uuid.MustParse(stringId)
	userId := r.Context().Value("userId").(string)
    //userId := uuid.MustParse(userIdString)

	note := db.GetNote(noteId, userId)

	b, err := json.Marshal(model.Response{Message: "Note retrieved", Note: note})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := db.GetAllNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.ResponseNotes{Data: notes}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func PostNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdNote, err := db.CreateNote(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.Response{Message: "Note was successfully created", Note: createdNote}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var noteToUpdate model.Note

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&noteToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
