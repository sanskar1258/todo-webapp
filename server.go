// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"todo/router"
	"todo/utility"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Initialize the database
	db := utility.InitDB()

	// Setup routes with controllers
	router.SetupRoutes(r, db)

	http.Handle("/", r)

	fmt.Println("Server is running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
