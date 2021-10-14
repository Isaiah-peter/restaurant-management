package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID				primitive.ObjectID	`json:"_id"`
	NumberOfGuests	*int				`json:"number_of_guests" validate:"required"`
	TableNumber		*int				`json:"table_number" validate:"required"`
	CreatedAt 		time.Time          	`json:"created_at"`
	UpdatedAt 		time.Time          	`json:"updated_at"`
	TableId    		string				`json:"table_id"`
}