// Issue 89
// Passing tainted data into db.ExecContext can
// result in sql injection.

package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	connStr := os.Getenv("DbConnStr")
	db, _ := sql.Open("postgres", connStr)
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")
	query := fmt.Sprintf("INSERT INTO users (username, password) VALUES (%s, %s)", username, password)
	ctx := context.Background()
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		// Handle the error
		fmt.Print("Error")
	}
	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		// Handle the error
		fmt.Print("Error")
	}
	fmt.Print(numRowsAffected)
}

func main() {
	http.HandleFunc("/", signupHandler)
	http.ListenAndServe(":8090", nil)
}
