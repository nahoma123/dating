package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
)

type Profile struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	ProfileID uuid.UUID     `bson:"profile_id" json:"profile_id,omitempty"`
	// Status is the status of the user.
	// It is set to active by default after successful registration.
	User   *User  `bson:"user" json:"user,omitempty"`
	Status string `json:"status,omitempty"`
	// ProfilePicture is the profile picture of the user.
	// It is set on a separate setProfilePicture endpoint.
	ProfilePicture string `json:"profile_picture,omitempty"`
	// CreatedAt is the time when the user is created.
	// It is automatically set when the user is created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u User) ValidateProfile() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Email, is.EmailFormat.Error("email is not valid")),
		validation.Field(&u.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
		validation.Field(&u.Password, validation.When(u.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
	)
}
