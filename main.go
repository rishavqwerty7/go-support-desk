package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/", handleRoutes)
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Home routes")
}
