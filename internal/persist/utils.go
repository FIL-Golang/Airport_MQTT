package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	config.LoadConfig()
}

func NewMongoClient() *mongo.Client {
	datasourceConfig := config.GetDatasourceConfig()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(datasourceConfig.Url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	return client
}
