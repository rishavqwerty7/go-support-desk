package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rishavqwerty7/go-support-desk/models"
	"github.com/rishavqwerty7/go-support-desk/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/tickets":
		handler.GetAllTicketsHandler(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/tickets/"):
		handler.GetTicketByIdHandler(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/tickets":
		handler.CreateTicketHandler(w, r)
	default:
		http.Error(w, "Methodd not allowed", http.StatusMethodNotAllowed)
	}

}

func (handler *TicketHandler) GetTicketByIdHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "invalide method", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/tickets/")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid ticket id", http.StatusBadRequest)
		return
	}

	res, err := handler.service.GetTicketById(r.Context(), objectId)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			http.Error(w, "Ticket not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
