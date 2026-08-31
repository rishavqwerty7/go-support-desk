package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/", handleRoutes)
	http.HandleFunc("/home", handleHomeRoute)
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Home routes")
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Home routes")
}
