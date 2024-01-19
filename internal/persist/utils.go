package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func init() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: file_recorder <config_file>")
		os.Exit(1)
	}
	config.LoadConfig(args[1])
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
