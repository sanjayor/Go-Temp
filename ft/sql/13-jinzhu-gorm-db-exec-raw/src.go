// Issue 89
// Passing tainted data into gorm.Db.Raw can
// result in sql injection.

package testdata

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	connStr := os.Getenv("DbConnStr")
	db, _ := gorm.Open("mysql", connStr)
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)

	var users []User
	// OpenRefactory Warning:
	// Possible SQL injection!
	// Path:
	//	File: src.go, Line: 26
	//		username := r.FormValue("username")
	//		Variable 'username' is assigned a tainted value from an external source.
	//	File: src.go, Line: 27
	//		password := r.FormValue("password")
	//		Variable 'password' is assigned a tainted value from an external source.
	//	File: src.go, Line: 28
	//		query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	//		Variable 'query' is assigned a tainted value which is passed through a function call.
	//	File: src.go, Line: 31
	//		db.Raw(query)
	//		Tainted information is used in a sink.
	db.Raw(query).Scan(&users)

	// Print the results
	for _, user := range users {
		fmt.Printf("id: %d, name: %s, email: %s\n", user.ID, user.Name, user.Email)
	}
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.ListenAndServe(":8090", nil)
}
