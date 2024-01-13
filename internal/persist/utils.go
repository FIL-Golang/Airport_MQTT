package persist

import (
	"Airport_MQTT/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	datasourceConfig := config.GetDatasourceConfig()

	clientOptions := options.Client()
	clientOptions.ApplyURI(datasourceConfig.Url)
	//clientOptions.SetAuth(options.Credential{ //TODO : implement option to connect to with credentials
	//	Username: datasourceConfig.Username,
	//	Password: datasourceConfig.Password,
	//})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}
