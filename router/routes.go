// router/routes.go
package router

import (
	"database/sql"
	"todo/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes routes with controllers
func SetupRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/", controllers.ServeHTML).Methods("GET")
	r.HandleFunc("/tasks", controllers.GetAllTasks(db)).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask(db)).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask(db)).Methods("DELETE")
}
