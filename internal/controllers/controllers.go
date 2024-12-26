package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	db "florent-haxhiu/todoInGo/internal/database"
	"florent-haxhiu/todoInGo/internal/model"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	string_id := chi.URLParam(r, "id")
	i, err := strconv.Atoi(string_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	note := db.GetNote(i)
	b, err := json.Marshal(model.Response{Message: "Note successfully retrieved", CreatedNote: note})
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

	createdNote, err := db.CreateNote(note, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.Response{Message: "Note was successfully created", CreatedNote: createdNote}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	// Update Note
}
