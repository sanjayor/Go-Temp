// Issue 89
// Passing tainted data into HostClient.Get can
// result in request forgery.

package main

import (
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	urlParam := r.URL.Query().Get("url")
	client := &fasthttp.HostClient{
		Addr: "example.com:80",
	}
	// Send the HTTP GET request
	statusCode, body, err := client.Get(nil, urlParam)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	// Print the response status code and body
	fmt.Printf("Status code: %d\n", statusCode)
	fmt.Printf("Response body: %s\n", body)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
