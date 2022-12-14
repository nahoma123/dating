package model

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Profile struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	ProfileID string        `bson:"profile_id" json:"profile_id,omitempty"`
	// Status is the status of the user.
	// It is set to active by default after successful registration.
	User   *User  `bson:"user" json:"user,omitempty"`
	Status string `json:"status,omitempty"`

	// about me info
	Title              string `bson:"title,omitempty" json:"title,omitempty"`
	AboutMe            string `bson:"about_me" json:"about_me,omitempty"`
	Age                int    `bson:"age" json:"age,omitempty"`
	Gender             string `bson:"gender,omitempty" json:"gender,omitempty"`
	RelationShipStatus string `bson:"relation_ship_status,omitempty" json:"relation_ship_status`
	Job                string `bson:"job,omitempty" json:"job,omitempty"`
	CompanyName        string `bson:"company_name,omitempty" json:"company_name,omitempty"`
	College            string `bson:"college,omitempty" json:"college,omitempty"`

	// ethnicity
	Country        string `bson:"country" json:"country,omitempty"`
	State          string `bson:"state" json:"state,omitempty"`
	Ethnicity      string `bson:"ethnicity" json:"ethnicity,omitempty"`
	LanguageSpoken string `bson:"language_spoken" json:"language,omitempty"`

	LookingFor []string `bson:"looking_for,omitempty" json:"looking_for"`

	// basic info
	Height      int    `bson:"height,omitempty" json:"height,omitempty"`
	Excercise   bool   `bson:"excercise,omitempty" json:"excercise,omitempty"`
	DrinkAlchol bool   `bson:"drink_alchol,omitempty" json:"drink_alchol,omitempty"`
	Children    int    `bson:"children,omitempty" json:"children,omitempty"`
	Religion    string `bson:"religion,omitempty" json:"religion,omitempty"`
	ZodiacSign  string `bson:"zodiac_sign,omitempty" json:"zodiac_sign"`
	Smoker      bool   `bson:"smoker,omitempty" json:"smoker,omitempty"`

	// ProfilePicture is the profile picture of the user.
	// It is set on a separate setProfilePicture endpoint.
	ProfilePicture []string `bson:"profile_picture,omitempty" json:"profile_pictures,omitempty"`
	// CreatedAt is the time when the user is created.
	// It is automatically set when the user is created.
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func (u Profile) ValidateProfile() error {
	if u.User == nil {
		return validation.Errors{
			"user": errors.New("is required"),
		}
	}
	return validation.ValidateStruct(&u, validation.Field(&u.User)) // validation.Field(&u.User.FirstName, validation.Required.Error("first name is required")),
	// validation.Field(&u.User.MiddleName, validation.Required.Error("middle name is required")),
	// validation.Field(&u.User.LastName, validation.Required.Error("last name is required")),
	// validation.Field(&u.User.Email, is.EmailFormat.Error("email is not valid")),
	// validation.Field(&u.User.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
	// validation.Field(&u.User.Password, validation.When(u.User.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
}
