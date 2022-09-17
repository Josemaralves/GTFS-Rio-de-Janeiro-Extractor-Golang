package repository

import (
	"context"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const StopsCollection = "stops"

type StopsRepository struct {
	collShapes *mongo.Collection
}

func NewStopRepository(db *mongo.Database) *StopsRepository {
	collStops := db.Collection(StopsCollection)

	return &StopsRepository{
		collShapes: collStops,
	}
}

func (r *StopsRepository) InsertStops(stops []models.Stop) {
	var bulk []mongo.WriteModel

	var i int
	for t, s := range stops {
		model := mongo.NewReplaceOneModel().SetFilter(bson.D{{"stopId", s.StopId}}).SetReplacement(s).SetUpsert(true)
		bulk = append(bulk, model)

		i++
		if i == 1000 || t == len(stops) {
			_, err := r.collShapes.BulkWrite(context.Background(), bulk)

			if err != nil {
				panic(err)
			}
			bulk = nil
			i = 0
		}

	}

}
