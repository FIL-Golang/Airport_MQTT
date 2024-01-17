package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func NewMongoClient() *mongo.Client {
	datasourceConfig := config.GetDatasourceConfig()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(datasourceConfig.Url).SetServerAPIOptions(serverAPI)

	slog.Info("Connecting to MongoDB...")

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	slog.Info("Connected to MongoDB")

	return client
}
