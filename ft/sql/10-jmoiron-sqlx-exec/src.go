// Issue 89
// Passing tainted data into sqlx.qStmt.Exec can
// result in sql injection.

package testdata

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func handler(w http.ResponseWriter, r *http.Request) {
	connStr := os.Getenv("DbConnStr")
	db, _ := sqlx.Connect("mysql", connStr)
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)

	// OpenRefactory Warning:
	// Possible SQL injection!
	// Path:
	//	File: src.go, Line: 21
	//		username := r.FormValue("username")
	//		Variable 'username' is assigned a tainted value from an external source.
	//	File: src.go, Line: 22
	//		password := r.FormValue("password")
	//		Variable 'password' is assigned a tainted value from an external source.
	//	File: src.go, Line: 23
	//		query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	//		Variable 'query' is assigned a tainted value which is passed through a function call.
	//	File: src.go, Line: 25
	//		db.Exec(query)
	//		Tainted information is used in a sink.
	db.Exec(query)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
