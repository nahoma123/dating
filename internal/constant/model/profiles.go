package model

import (
	"time"

	"github.com/google/uuid"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Profile struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID uuid.UUID `json:"id,omitempty"`
	// FirstName is the first name of the user.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName is the middle name of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName is the last name of the user.
	LastName string `json:"last_name,omitempty"`
	// Email is the email of the user.
	Email string `json:"email,omitempty"`
	// Phone is the phone of the user.
	Phone string `json:"phone,omitempty"`
	// Password is the password of the user.
	// It is only used for logging in with email
	Password string `json:"password,omitempty"`
	// UserName is the username of the user.
	// It is currently of no use
	UserName string `json:"user_name,omitempty"`
	// Gender is the gender of the user.
	Gender string `json:"gender,omitempty"`
	// Status is the status of the user.
	// It is set to active by default after successful registration.
	Status string `json:"status,omitempty"`
	// ProfilePicture is the profile picture of the user.
	// It is set on a separate setProfilePicture endpoint.
	ProfilePicture string `json:"profile_picture,omitempty"`
	// CreatedAt is the time when the user is created.
	// It is automatically set when the user is created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// OTP is the one time password of the user.
	OTP string `json:"otp,omitempty"`
}

func (u User) ValidateProfile() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Email, is.EmailFormat.Error("email is not valid")),
		validation.Field(&u.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
		validation.Field(&u.Password, validation.When(u.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
		validation.Field(&u.OTP, validation.Required.Error("otp is required"), validation.Length(6, 6).Error("otp must be 6 characters")),
	)
}
