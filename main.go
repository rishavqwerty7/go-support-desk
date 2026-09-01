package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server running", err)
	}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Support Desk Api Healthy"))
}
