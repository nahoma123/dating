package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CountryId string             `bson:"country_id,omitempty" json:"country_id"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type State struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StateId   string             `bson:"state_id,omitempty" json:"state_id"`
	CountryId string             `bson:"country_id,omitempty" json:"country_id"`
	Name      string             `json:"name" bson:"name,omitempty" validate:"required"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Ethnicity struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EthnicityId string             `bson:"ethnicity_id,omitempty" json:"ethnicity_id"`
	CountryId   string             `bson:"country_id,omitempty" json:"country_id"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
