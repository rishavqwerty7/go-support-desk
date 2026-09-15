package services

import (
	"context"
	"time"

	"github.com/rishavqwerty7/go-support-desk/models"
	"github.com/rishavqwerty7/go-support-desk/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketService struct {
	repository *repositories.TicketRepository
}

func NewTicketService(repository *repositories.TicketRepository) *TicketService {
	return &TicketService{
		repository: repository,
	}
}

func (service *TicketService) CreateTicket(
	ctx context.Context,
	request models.CreateTicketRequest,
	customerId primitive.ObjectID,
) error {

	ticket := models.Ticket{
		Title:       request.Title,
		Description: request.Description,
		Priority:    request.Priority,
		CustomerId:  customerId,
		Status:      "open",
		CreatedAt:   time.Now(),
	}

	return service.repository.Create(ctx, &ticket)

}

func (service *TicketService) GetAllTickets(ctx context.Context) ([]models.Ticket, error) {

	tickets, err := service.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return tickets, nil

}

func (service *TicketService) GetTicketById(ctx context.Context, id primitive.ObjectID) (*models.Ticket, error) {

	ticket, err := service.repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (service *TicketService) UpdateTicket(ctx context.Context, id primitive.ObjectID, update bson.M) error {

	err := service.repository.Update(ctx, id, update)

	if err != nil {
		return err
	}

	return nil

}
