// Issue 176
// Passing tainted data into http.Client.Do can
// result in request forgery.

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("failed to create request")
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		panic("failed to send request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(w, "Error fetching URL: %v", resp.Status)
		return
	}
	fmt.Fprintf(w, "Success")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
