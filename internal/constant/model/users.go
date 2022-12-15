package model

import (
	"fmt"

	"github.com/dongri/phonenumber"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
)

type User struct {
	ID *int `bson:"_id,omitempty" json:"id"`
	// FirstName is the first name of the user.
	UserID    uuid.UUID `bson:"user_id,omitempty" json:"user_id"`
	FirstName string    `bson:"first_name,omitempty"  json:"first_name,omitempty"`
	// MiddleName is the middle name of the user.
	MiddleName string `bson:"middle_name,omitempty" json:"middle_name,omitempty"`
	// LastName is the last name of the user.
	LastName string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	// Email is the email of the user.
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	// Phone is the phone of the user.
	Phone string `bson:"phone,omitempty" json:"phone,omitempty"`
	// Password is the password of the user.
	// It is only used for logging in with email
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	// UserName is the username of the user.
	// It is currently of no use
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
	// Gender is the gender of the user.
	Gender string `bson:"gender,omitempty" json:"gender,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Email, is.EmailFormat.Error("email is not valid")),
		validation.Field(&u.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
		validation.Field(&u.Password, validation.When(u.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
	)
}

func validatePhone(phone interface{}) error {
	str := phonenumber.Parse(fmt.Sprintf("%v", phone), "ET")
	if str == "" {
		return fmt.Errorf("invalid phone number")
	}
	return nil
}
