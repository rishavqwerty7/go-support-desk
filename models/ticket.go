package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	Priority    string             `json:"priority" bson:"priority"`
	CustomerId  primitive.ObjectID `json:"customerId" bson:"customerId"`
	AssignedTo  primitive.ObjectID `json:"assignedTo" bson:"assignedTo"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description" bson:"description"`
	Priority    string `json:"priority" bson:"priority"`
}
