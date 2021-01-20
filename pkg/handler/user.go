package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Ariesfall/simple-auth-go/pkg/data"
	"github.com/jmoiron/sqlx"
)

func UserGet(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			writeErrorStatus(w, errors.New("User id not found"))
			return
		}

		idint, _ := strconv.Atoi(id)
		in := &data.Usrs{ID: idint}
		res, err := data.UserGet(db, in)
		if err != nil {
			writeErrorStatus(w, err)
			return
		}

		writeJsonStatus(w, res)
	}
}

func UserCreate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &data.Usrs{}
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		err = data.UserCreate(db, in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		w.Write([]byte("ok"))
	}
}

func UserUpdate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &data.Usrs{}
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		err = data.UserUpdate(db, in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		w.Write([]byte("ok"))
	}
}

func UserDelete(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &data.Usrs{}
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		err = data.UserDelete(db, in)
		if err != nil {
			writeErrorStatus(w, err)
		}

		w.Write([]byte("ok"))
	}
}

func writeJsonStatus(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeErrorStatus(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
