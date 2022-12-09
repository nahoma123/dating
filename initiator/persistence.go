package initiator

import (
	"dating/platform/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Persistence struct {
	// TODO implement
}

func InitPersistence(db *mongo.Client, log logger.Logger) Persistence {
	return Persistence{}
}
