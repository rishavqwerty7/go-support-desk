package repositories

import (
	"context"
	"fmt"

	"github.com/rishavqwerty7/go-support-desk/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketRepository struct {
	collection *mongo.Collection
}

func NewTicketRepository(collection *mongo.Collection) *TicketRepository {
	return &TicketRepository{
		collection: collection,
	}
}

func (repo *TicketRepository) Create(ctx context.Context,
	ticket *models.Ticket,
) error {
	_, err := repo.collection.InsertOne(ctx, ticket)
	return err
}

func (repo *TicketRepository) GetAll(ctx context.Context) 
([]models.Ticket, error) {

	fmt.Println("hi")
}
