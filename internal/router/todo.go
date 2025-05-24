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
	userIdString, ok := r.Context().Value("userId").(string)

	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	note := db.GetNote(noteId, userIdString)

	b, err := json.Marshal(model.Response{Message: "Note retrieved", Note: note})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	userIdString, ok := r.Context().Value(model.UserId("userId")).(string)

	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	notes, err := db.GetAllNotes(userIdString)
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

	k := model.UserId("userId")

	userIdString, ok := r.Context().Value(k).(string)

	if !ok {
		err_resp, _ := json.Marshal(model.ErrorResponse{Message: "User ID was not found", Status: http.StatusUnprocessableEntity})
		http.Error(w, string(err_resp), http.StatusUnprocessableEntity)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdNote, err := db.CreateNote(note, userIdString)
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
