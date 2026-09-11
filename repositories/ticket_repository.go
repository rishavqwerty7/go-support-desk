package repositories

import (
	"context"

	"github.com/rishavqwerty7/go-support-desk/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *TicketRepository) GetAll(
	ctx context.Context,
) ([]models.Ticket, error) {

	cursor, err := repo.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var tickets []models.Ticket

	for cursor.Next(ctx) {

		var ticket models.Ticket

		err = cursor.Decode(&ticket)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tickets, nil

}

func (repo *TicketRepository) GetById(ctx context.Context,
	id primitive.ObjectID) (*models.Ticket, error) {

	var ticket models.Ticket

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&ticket)

	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (repo *TicketRepository) Update(ctx context.Context,
	id primitive.ObjectID,
	update bson.M,
) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": update},
	)

	if err != nil {
		return err
	}

	return nil

}

func (repo *TicketRepository) Delete(ctx context.Context, id primitive.ObjectID) error {

	_, err := repo.collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	if err != nil {
		return err
	}

	return nil

}
