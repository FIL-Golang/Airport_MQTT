package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	conf := config.LoadConfig()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Datasource.Url))
	if err != nil {
		panic(err)
	}
	return client
}
