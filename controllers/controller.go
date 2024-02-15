// controllers/controller.go
package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"
	"todo/models"

	"github.com/gorilla/mux"
)

// ServeHTML serves the HTML file to the browser
func ServeHTML(w http.ResponseWriter, r *http.Request) {
	htmlPath, err := filepath.Abs("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, htmlPath)
}

// GetAllTasks handles the request to retrieve all tasks
func GetAllTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := models.GetAllTasks(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

// CreateTask handles the request to create a new task
func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err = models.CreateTask(db, &task)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// DeleteTask handles the request to delete a task
func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskID := params["id"]

		err := models.DeleteTask(db, taskID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
