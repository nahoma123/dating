package constant

import (
	"context"
	"dating/internal/constant/errors"
	"dating/platform/logger"
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Filters
 1. by gte for date, by gte for age [gte]
 2. equals [=]
 3. contains [contains]
 4. not equals [!=]
*/
func GetResults(cxt context.Context, db *mongo.Database, collectionName string, filterPagination *FilterPagination, results *[]bson.M, geoFilter bson.M) error {
	// Get a reference to the collection
	collection := db.Collection(collectionName)

	// Create slice to hold filter stage documents
	filters := make([]bson.M, 0)

	// Iterate over filters and create filter stage documents
	for _, f := range filterPagination.Filters {
		var filter bson.M

		if f.Operator == "gte" {
			if f.Field == "created_at" || f.Field == "updated_at" {
				t, err := time.Parse("2006-01-02T15:04:05", f.Value)
				if err != nil {
					fmt.Printf("Error parsing time: %v", err)
					return errors.ErrAcessError.Wrap(err, "wrong time format")
				}
				filter = bson.M{f.Field: bson.M{"$gte": t}}
			} else {
				filter = bson.M{"created_at": bson.M{"$gte": f.Value}}
			}
		} else if f.Operator == "=" {
			filter = bson.M{f.Field: f.Value}
		} else if f.Operator == "contains" {
			filter = bson.M{f.Field: bson.M{"$regex": primitive.Regex{Pattern: f.Value, Options: "i"}}}
		} else if f.Operator == "!=" {
			filter = bson.M{f.Field: bson.M{"$ne": f.Value}}
		} else {
			// Handle other operator types
		}
		filters = append(filters, filter)
	}

	// Create pipeline with filter stage if there are filters
	var pipeline []bson.M

	pipeline = []bson.M{}
	if len(geoFilter) > 0 {
		pipeline = append(pipeline, geoFilter)
	}

	if len(filters) > 0 {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"$and": filters}})
	}

	// Add sort stage to pipeline if there are sort fields
	if len(filterPagination.Pagination.Sort) > 0 {
		var sortFields []bson.M
		for field, direction := range filterPagination.Pagination.Sort {
			if direction == "asc" {
				sortFields = append(sortFields, bson.M{field: 1})
			} else if direction == "desc" {
				sortFields = append(sortFields, bson.M{field: -1})
			} else {
				// Handle other sort directions
			}
		}
		pipeline = append(pipeline, bson.M{"$sort": bson.M{"$and": sortFields}})
	}

	// Add pagination stages to pipeline
	skip := (filterPagination.Pagination.Page - 1) * filterPagination.Pagination.PerPage
	pipeline = append(pipeline, bson.M{"$skip": skip}, bson.M{"$limit": filterPagination.Pagination.PerPage})

	// Execute pipeline and retrieve results
	ctx, cancel := context.WithTimeout(cxt, 30*time.Second)
	defer cancel()

	fmt.Printf("Slice: %#v\n", pipeline)

	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}
	defer cur.Close(ctx)

	// Iterate over results and append them to the results slice
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)

		}
		*results = append(*results, result)
	}

	if err := cur.Err(); err != nil {
		return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}

	// Get the total count of documents that match the filter
	totalCount := int64(0)
	if len(filters) != 0 {
		totalCount, err = collection.CountDocuments(ctx, bson.M{"$and": filters})
		if err != nil {
			logger.Log().Error(ctx, fmt.Sprintf("error on db count error %s", err.Error()))
			return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)

		}
	} else {
		totalCount, err = collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			logger.Log().Error(ctx, fmt.Sprintf("error on db count error %s", err.Error()))
			return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)

		}
	}

	filterPagination.TotalCount = totalCount
	filterPagination.TotalPages = int(math.Ceil(float64(totalCount) / float64(filterPagination.Pagination.PerPage)))

	return nil
}
