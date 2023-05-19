// Issue 89
// Avoid passing hard coded ip into net.Dial

package main

import (
	"fmt"
	"net/http"
)

func main() {
	connect("127.0.0.1:8080")
}

func connect(ip string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(ip, nil)
}
