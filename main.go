package main

import (
	"fmt"
	"net/http"
)

// go desk main func
func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server running", err)
	}

}

// A health endpoint is a small API endpoint that allows another system to ask your application whether it is working properly.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Support Desk Api Healthy"))
}

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home route"))
}
