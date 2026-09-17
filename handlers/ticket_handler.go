package handlers

import "github.com/rishavqwerty7/go-support-desk/services"

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {

	return &TicketHandler{
		service: service,
	}
}
