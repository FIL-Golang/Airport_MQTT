package persist

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/config/types"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	conf := config.LoadConfig(&types.ConfigFile{}, "config.yaml").(*types.ConfigFile)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Datasource.Url))
	if err != nil {
		panic(err)
	}
	return client
}
