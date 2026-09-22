package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rishavqwerty7/go-support-desk/models"
	"github.com/rishavqwerty7/go-support-desk/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {

	return &TicketHandler{
		service: service,
	}
}

func (handler *TicketHandler) CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Status method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var ticket models.CreateTicketRequest

	err := json.NewDecoder(r.Body).Decode(&ticket)

	if err != nil {
		http.Error(w, "Inavlid request body", http.StatusBadRequest)
	}

	customerId := primitive.NewObjectID()
	err = handler.service.CreateTicket(r.Context(), ticket, customerId)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("ticket created successfully"))

}

func (handler *TicketHandler) GetAllTicketsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid status Method", http.StatusMethodNotAllowed)
		return
	}

	tickets, err := handler.service.GetAllTickets(r.Context())

	if err != nil {
		http.Error(w, "Internal serval error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tickets)

}

func (handler *TicketHandler) TicketHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handler.GetAllTicketsHandler(w, r)
	case http.MethodPost:
		handler.CreateTicketHandler(w, r)
	default:
		http.Error(w, "Methodd not allowed", http.StatusMethodNotAllowed)
	}

}
