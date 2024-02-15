// utility/db.go
package utility

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "reak:reak@tcp(localhost:3306)/todo")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
