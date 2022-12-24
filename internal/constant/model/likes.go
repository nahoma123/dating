package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SelectProfileId struct {
	ProfileId string `json:"profile_id,omitempty" bson:"-"`
}

type Likes struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID             string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	DisLikedProfileIDs []string           `bson:"disliked_profile_ids,omitempty" json:"disliked_profile_ids,omitempty"`
	LikedProfileIDs    []string           `bson:"liked_profile_ids,omitempty" json:"liked_profile_ids,omitempty"`
	CreatedAt          time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt          time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Favorites struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LikedProfileIDs []string           `bson:"liked_profile_ids,omitempty" json:"liked_profile_ids,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
