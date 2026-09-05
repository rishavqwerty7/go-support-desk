package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishavqwerty7/go-support-desk/database"
)

// go desk main func
func main() {
	err := database.ConnectMongoDB()

	if err != nil {
		log.Fatal("mongo db connection failed")
	}

	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server failed", err)
	}
}

// A health endpoint is a small API endpoint that allows another system to ask your application whether it is working properly.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Support Desk Api Healthy"))
}

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home route"))
}
