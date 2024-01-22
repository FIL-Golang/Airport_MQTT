package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func init() {
	config.LoadConfig()
}

func NewMongoClient() *mongo.Client {
	datasourceConfig := config.GetDatasourceConfig()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(datasourceConfig.Url).SetServerAPIOptions(serverAPI)

	slog.Info("Connecting to MongoDB")
	client, err := mongo.Connect(context.TODO(), opts)
	slog.Info("Connected to MongoDB")

	if err != nil {
		panic(err)
	}
	return client
}
