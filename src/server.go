package mains

import (
    "fmt"
    "net/http"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.PrintLn("Hello World")
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.PrintLn("Hello..")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetRequest(w, r)
		} else if r.method == "POST" {
			handlePostRequest(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServer(":8080", nil)
}