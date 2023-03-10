package initiator

import (
	"context"
	"dating/internal/storage"
	"dating/internal/storage/persistence"
	"dating/platform/logger"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Persistence struct {
	// TODO implement
	Profile storage.ProfileStorage
	Mesc    storage.MescStorage
}

func CreateIndexes(log logger.Logger, db *mongo.Database) {
	log.Info(context.Background(), "create indexes")

	_, err := db.Collection(string(storage.Profile)).Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys: bson.D{
				bson.E{Key: "_id", Value: 1},
				bson.E{Key: "email", Value: 1},
				bson.E{Key: "user_name", Value: 1},
				bson.E{Key: "phone", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Debug(context.Background(), fmt.Sprint("create indexes error: ", err.Error()))
	}

	_, err = db.Collection(string(storage.Country)).Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys: bson.D{
				bson.E{Key: "name", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Debug(context.Background(), fmt.Sprint("create indexes error: ", err.Error()))
	}
}

func InitPersistence(db *mongo.Database, log logger.Logger) Persistence {

	profileStorage := persistence.InitProfileDB(db)
	mescStorage := persistence.InitMescDB(db)
	return Persistence{
		Profile: profileStorage,
		Mesc:    mescStorage,
	}
}
