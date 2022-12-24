package model

type Location struct {
	Type        string    `bson:"type,omitempty" json:"type,omitempty"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}
