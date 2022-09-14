package repository

import (
	"context"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = ""
const Database = ""
const Collection = ""

type RouteRepository struct {
	collRoutes *mongo.Collection
}

func New() *RouteRepository {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}

	collRoutes := client.Database(Database).Collection(Collection)

	return &RouteRepository{
		collRoutes: collRoutes,
	}
}

func (r *RouteRepository) InsertRoutes(routes []models.Route) {
	var bulk []interface{}

	for _, r := range routes {
		bulk = append(bulk, r)
	}

	_, err := r.collRoutes.InsertMany(context.Background(), bulk)

	if err != nil {
		panic(err)
	}

}
