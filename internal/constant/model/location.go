package model

type Location struct {
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}
