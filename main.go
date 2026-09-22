package main

import (
	"log"
	"net/http"

	"github.com/rishavqwerty7/go-support-desk/database"
	"github.com/rishavqwerty7/go-support-desk/handlers"
	"github.com/rishavqwerty7/go-support-desk/repositories"
	"github.com/rishavqwerty7/go-support-desk/services"
)

// go desk main func
func main() {
	err := database.ConnectMongoDB()

	if err != nil {
		log.Fatal("Mongo DB Connection Failed")
	}

	collection := database.Client.Database("go_support_desk").Collection("tickets")

	ticketRepository := repositories.NewTicketRepository(collection)

	ticketService := services.NewTicketService(ticketRepository)

	ticketHandler := handlers.NewTicketHandler(ticketService)

	http.HandleFunc("/tickets", ticketHandler.TicketHandler)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server failed", err)
	}
}

// // A health endpoint is a small API endpoint that allows another system to ask your application whether it is working properly.
// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Support Desk Api Healthy"))
// }

// func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Home route"))
// }
