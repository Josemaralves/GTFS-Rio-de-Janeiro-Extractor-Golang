package repository

import (
	"context"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const RoutesCollection = "lines"

type RouteRepository struct {
	collRoutes *mongo.Collection
}

func NewRouteRepository(db *mongo.Database) *RouteRepository {
	collRoutes := db.Collection(RoutesCollection)

	return &RouteRepository{
		collRoutes: collRoutes,
	}
}

func (r *RouteRepository) InsertRoutes(routes []models.Lines) {
	var bulk []mongo.WriteModel

	for _, r := range routes {
		model := mongo.NewReplaceOneModel().SetFilter(bson.D{{"idRoute", r.IdRoute}}).SetReplacement(r).SetUpsert(true)
		bulk = append(bulk, model)
	}

	_, err := r.collRoutes.BulkWrite(context.Background(), bulk)

	if err != nil {
		panic(err)
	}

}
