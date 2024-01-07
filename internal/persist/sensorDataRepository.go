package persist

import (
	"Airport_MQTT/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

const databaseName = "airport_sensors"
const collectionName = "sensor_data"

type SensorDataRepository interface {
	Store(data model.SensorData) (savedData model.SensorData, err error)
}

type sensorDataMongoRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewSensorDataRepository() SensorDataRepository {
	client := NewMongoClient()
	collection := client.Database(databaseName).Collection(collectionName)
	return &sensorDataMongoRepository{
		collection: collection,
		ctx:        context.Background(),
	}
}

func (r *sensorDataMongoRepository) Store(data model.SensorData) (savedData model.SensorData, err error) {
	_, err = r.collection.InsertOne(r.ctx, data)
	if err != nil {
		return data, err
	}
	return data, nil
}
