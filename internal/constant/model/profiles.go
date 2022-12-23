package model

import (
	errorsconstats "dating/internal/constant/errors"
	"errors"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProfileID string             `bson:"profile_id,omitempty" json:"profile_id,omitempty"`
	// Status is the status of the user.
	// It is set to active by default after succesObjectIDs must be exactly 12 bytes long (got 0)sful registration.
	User   *User  `bson:"user,omitempty" json:"user,omitempty"`
	Status string `json:"status,omitempty"`

	// about me info
	IsSpotLight        bool   `bson:"is_spot_light,omitempty" json:"is_spot_light"`
	Title              string `bson:"title,omitempty" json:"title,omitempty"`
	AboutMe            string `bson:"about_me,omitempty" json:"about_me,omitempty"`
	Age                int    `bson:"age,omitempty" json:"age,omitempty"`
	Gender             string `bson:"gender,omitempty" json:"gender,omitempty"`
	RelationShipStatus string `bson:"relation_ship_status,omitempty" json:"relation_ship_status,omitempty"`
	Job                string `bson:"job,omitempty" json:"job,omitempty"`
	CompanyName        string `bson:"company_name,omitempty" json:"company_name,omitempty"`
	College            string `bson:"college,omitempty" json:"college,omitempty"`

	// ethnicity
	Country        string `bson:"country" json:"country,omitempty"`
	State          string `bson:"state" json:"state,omitempty"`
	Ethnicity      string `bson:"ethnicity" json:"ethnicity,omitempty"`
	LanguageSpoken string `bson:"language_spoken" json:"language_spoken,omitempty"`

	LookingFor []string `bson:"looking_for,omitempty" json:"looking_for"`

	// basic info
	Height      int    `bson:"height,omitempty" json:"height,omitempty"`
	Excercise   bool   `bson:"excercise,omitempty" json:"excercise,omitempty"`
	DrinkAlchol bool   `bson:"drink_alchol,omitempty" json:"drink_alchol,omitempty"`
	Children    int    `bson:"children,omitempty" json:"children"`
	Religion    string `bson:"religion,omitempty" json:"religion,omitempty"`
	ZodiacSign  string `bson:"zodiac_sign,omitempty" json:"zodiac_sign"`
	Smoker      bool   `bson:"smoker,omitempty" json:"smoker,omitempty"`

	// ProfilePicture is the profile picture of the user.
	// It is set on a separate setProfilePicture endpoint.
	ProfilePicture []string `bson:"profile_picture,omitempty" json:"profile_pictures,omitempty"`
	// CreatedAt is the time when the user is created.
	// It is automatically set when the user is created.

	Distance float64   `bson:"distance" json:"distance"`
	Location []float64 `bson:"location,omitempty" json:"location,omitempty"`

	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func (u Profile) ValidateRegisterProfile() error {
	if u.User == nil {
		return validation.Errors{
			"user": errors.New("is required"),
		}
	}

	return validation.ValidateStruct(&u, validation.Field(&u.User), // validation.Field(&u.User.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.AboutMe, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "about_me"))),
		validation.Field(&u.Age, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "age"))),
		// validation.Field(&u.Children, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "children"))),
		validation.Field(&u.College, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "college"))),
		validation.Field(&u.CompanyName, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "company_name"))),
		validation.Field(&u.Country, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "country"))),
		// validation.Field(&u.DrinkAlchol, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "drink_alchol"))),
		validation.Field(&u.Ethnicity, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "ethnicity"))),
		validation.Field(&u.Excercise, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "excercise"))),

		validation.Field(&u.Job, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "job"))),
		validation.Field(&u.Gender, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "gender"))),
		validation.Field(&u.Height, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "height"))),
		validation.Field(&u.LanguageSpoken, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "language_spoken"))),
		validation.Field(&u.LookingFor, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "looking_for"))),
		validation.Field(&u.RelationShipStatus, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "relation_ship_status"))),
		validation.Field(&u.RelationShipStatus, validation.In("Single", "InRelationShip")),
		validation.Field(&u.Religion, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "religion"))),
		// validation.Field(&u.Smoker, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "smoker"))),
		validation.Field(&u.State, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "state"))),
		validation.Field(&u.Title, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "title"))),
		validation.Field(&u.ZodiacSign, validation.Required.Error(fmt.Sprintf(errorsconstats.FeildIsRequired, "zodiac_sign"))),
	)
	// validation.Field(&u.User.Phone, validation.Required.Error("phone is required"), validation.By(validatePhone)),
	// validation.Field(&u.User.Password, validation.When(u.User.Email != "", validation.Required.Error("password is required"), validation.Length(6, 32).Error("password must be between 6 and 32 characters"))),
}

func (u Profile) ValidateUpdateProfile() error {
	return validation.ValidateStruct(&u, validation.Field(&u.RelationShipStatus, validation.In("Single", "InRelationShip"))) // validation.Field(&u.User.FirstName, validation.Required.Error("first name is required")),
}
