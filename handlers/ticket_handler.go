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

func (handler *TicketHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
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

}
