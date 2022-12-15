package model

import (
	"time"
)

type Country struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID        *int      `bson:"_id,omitempty" json:"_id,omitempty"`
	CountryId string    `bson:"country_id,omitempty" json:"country_id"`
	Name      string    `json:"name" bson:"name" validate:"required"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type State struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID        *int      `bson:"_id,omitempty" json:"_id,omitempty"`
	CountryId string    `bson:"country_id,omitempty" json:"country_id"`
	Name      string    `json:"name" bson:"name" validate:"required"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Ethnicity struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID        *int      `bson:"_id,omitempty" json:"_id,omitempty"`
	CountryId string    `bson:"country_id,omitempty" json:"country_id"`
	Name      string    `json:"name" bson:"name" validate:"required"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
