package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/DavinY/go-restful-gin/config"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

//Init..
func Init() (err error) {
	ctx, cancel := NewMongoContext()
	defer cancel()

	c := config.Get()

	clientOptions := options.Client().ApplyURI(c.Database.Uri + c.Port.HTTPServer)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return errors.Wrapf(err, "error connecting to mongodb")
	}

	log.Println("Connected to database")
	mongoClient = client

	return
}

//GetDB..
func GetDB() *mongo.Database {
	cfg := config.Get()
	database := mongoClient.Database(cfg.Database.Name)
	return database
}

func GetCollection(db string, coll string) *mongo.Collection {
	// cfg := config.Get()
	collection := mongoClient.Database(db).Collection(coll)
	return collection
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
