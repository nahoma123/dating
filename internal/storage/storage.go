package storage

import (
	"context"
	"dating/internal/constant/model"
)

type DatabaseCollection string

// database collection constants
const (
	Profile DatabaseCollection = "profiles"
)

type ProfileStorage interface {
	Create(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Update(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Get(ctx context.Context, id string) (*model.Profile, error)
}
