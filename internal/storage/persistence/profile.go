package persistence

import (
	"context"
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"
	"dating/platform/logger"
	"fmt"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type profile struct {
	db *mongo.Database
}

func InitProfileDB(db *mongo.Database) storage.ProfileStorage {
	return &profile{
		db: db,
	}
}

func (p *profile) Create(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )
	id, _ := uuid.NewV4()
	profile.ProfileID = id.String()

	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	profile.Status = constant.Active
	_, err := p.db.Collection(string(storage.Profile)).InsertOne(ctx, profile)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
	}

	return profile, err
}

func (p *profile) Update(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )

	profile.UpdatedAt = time.Now()
	pByte, err := bson.Marshal(profile)
	if err != nil {
		return nil, err
	}

	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return nil, err
	}

	_, err = p.db.Collection(string(storage.Profile)).UpdateOne(ctx, bson.M{"profile_id": profile.ProfileID}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
	}

	return profile, err
}

func (p *profile) Get(ctx context.Context, id string) (*model.Profile, error) {
	profile := &model.Profile{}
	err := p.db.Collection(string(storage.Profile)).FindOne(ctx, bson.M{"profile_id": id}).Decode(profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNoRecordFound.New(errors.RecordNotfound)
		}
		return nil, err
	}
	return profile, nil
}

func (p *profile) GetLike(ctx context.Context, user_id string) (*model.Likes, error) {
	profile := &model.Likes{}
	err := p.db.Collection(string(storage.Like)).FindOne(ctx, bson.M{"user_id": user_id}).Decode(profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNoRecordFound.New(errors.RecordNotfound)
		}
		return nil, err
	}
	return profile, nil
}

func (p *profile) GetCustomers(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Profile, error) {
	var results []bson.M

	distance := constant.ExtractValueFromFilter(filterPagination, "distance")
	var distanceInt int
	var err error
	if distance != "" {
		distanceInt, err = strconv.Atoi(distance)
		if err != nil {
			return nil, errors.ErrReadError.Wrap(err, "invalid location value")
		}

		// remove the location filter
		constant.DeleteFilter(filterPagination, "distance")

		profileId := constant.ExtractValueFromFilter(filterPagination, "profile_id")
		if profileId == "" {
			return nil, errors.ErrReadError.New("invalid user id")
		}

		pf, err := p.Get(ctx, profileId)
		if err != nil {
			return nil, err
		}

		location := []float64{}

		fmt.Println("%s", pf)
		if pf.Location != nil {
			location = pf.Location.Coordinates
		}

		if len(location) == 2 {
			filter := constant.LocationFilter(location, distanceInt)

			err := constant.GetResults(ctx, p.db, string(storage.Profile), filterPagination, &results, filter)
			if err != nil {
				return nil, err
			}
		}
	} else {
		profileId := constant.ExtractValueFromFilter(filterPagination, "profile_id")
		if profileId == "" {
			return nil, errors.ErrReadError.New("invalid user id")
		}

		err := constant.GetResults(ctx, p.db, string(storage.Profile), filterPagination, &results, nil)
		if err != nil {
			return nil, err
		}
	}

	var profiles []model.Profile
	for _, result := range results {
		var profile model.Profile
		b, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(b, &profile)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (p *profile) LikeProfile(ctx context.Context, userID string, profileID string) error {
	// create filter to find the user being liked
	filter := bson.M{"profile_id": profileID}

	// use FindOne to check if the user being liked exists
	var result bson.M
	err := p.db.Collection(string(storage.Profile)).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// return an error if the user being liked does not exist
			return fmt.Errorf("user with ID %s does not exist", profileID)
		}
		return err
	}

	// create filter to find the user's Likes document
	filter = bson.M{"user_id": userID}

	// create update to add the profile ID to the LikedProfileIDs array
	update := bson.M{"$addToSet": bson.M{"liked_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err = p.db.Collection(string(storage.Like)).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) UnlikeProfile(ctx context.Context, userID string, profileID string) error {
	err := p.RemoveDislikeProfile(ctx, userID, profileID)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err, err.Error())
	}

	// create filter to find the user's Likes document
	filter := bson.M{"user_id": userID}

	// create update to remove the profile ID from the LikedProfileIDs array
	update := bson.M{"$pull": bson.M{"liked_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err = p.db.Collection(string(storage.Like)).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) MakeFavorite(ctx context.Context, userID string, profileID string) error {
	// create filter to find the user being favorite
	filter := bson.M{"profile_id": profileID}

	// use FindOne to check if the user being favorite exists
	var result bson.M
	err := p.db.Collection(string(storage.Profile)).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// return an error if the user being favorite does not exist
			return fmt.Errorf("user with ID %s does not exist", profileID)
		}
		return err
	}

	// create filter to find the user's Favorites document
	filter = bson.M{"user_id": userID}

	// create update to add the profile ID to the favoriteProfileIDs array
	update := bson.M{"$addToSet": bson.M{"favorite_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err = p.db.Collection(string(storage.Favorite)).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) RemoveFavorite(ctx context.Context, userID string, profileID string) error {
	// create filter to find the user's Favorite document
	filter := bson.M{"user_id": userID}

	// create update to remove the profile ID from the LikedProfileIDs array
	update := bson.M{"$pull": bson.M{"favorite_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err := p.db.Collection(string(storage.Favorite)).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) DisLikeProfile(ctx context.Context, userID string, profileID string) error {
	// create filter to find the user being disliked
	filter := bson.M{"profile_id": profileID}

	err := p.UnlikeProfile(ctx, userID, profileID)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err, err.Error())
	}
	// use FindOne to check if the user being disliked exists
	var result bson.M
	err = p.db.Collection(string(storage.Profile)).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// return an error if the user being disliked does not exist
			return fmt.Errorf("user with ID %s does not exist", profileID)
		}
		return err
	}

	// create filter to find the user's Dislikes document
	filter = bson.M{"user_id": userID}

	// create update to add the profile ID to the DislikedProfileIDs array
	update := bson.M{"$addToSet": bson.M{"disliked_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err = p.db.Collection(string(storage.Like)).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) RemoveDislikeProfile(ctx context.Context, userID string, profileID string) error {
	// create filter to find the user's Dislikes document
	filter := bson.M{"user_id": userID}

	// create update to remove the profile ID from the DislikedProfileIDs array
	update := bson.M{"$pull": bson.M{"disliked_profile_ids": profileID}}

	// use UpdateOne to apply the update
	_, err := p.db.Collection(string(storage.Like)).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (p *profile) GetRecommendations(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Profile, error) {
	var results []bson.M

	distance := constant.ExtractValueFromFilter(filterPagination, "distance")
	var distanceInt int
	var err error

	distanceInt, err = strconv.Atoi(distance)
	if err != nil {
		return nil, errors.ErrReadError.Wrap(err, "invalid location value")
	}

	// remove the location filter
	constant.DeleteFilter(filterPagination, "distance")

	profileId := constant.ExtractValueFromFilter(filterPagination, "profile_id")
	if profileId == "" {
		return nil, errors.ErrReadError.New("invalid user id")
	}

	pf, err := p.Get(ctx, profileId)
	if err != nil {
		return nil, err
	}

	location := []float64{}

	if pf.Location != nil {
		location = pf.Location.Coordinates
	}

	if len(location) == 2 {
		filter := constant.LocationFilter(location, distanceInt)
		like, _ := p.GetLike(ctx, profileId)
		err := constant.GetRecommendationsDb(ctx, p.db, string(storage.Profile), filterPagination, &results, filter, like.DisLikedProfileIDs)
		if err != nil {
			return nil, err
		}
	}

	var profiles []model.Profile
	for _, result := range results {
		var profile model.Profile
		b, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(b, &profile)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
