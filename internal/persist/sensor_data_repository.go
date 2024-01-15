package persist

import (
	"Airport_MQTT/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "airport_sensors"
	COLLECTION = "sensor_data"
)

type Filter struct {
	SensorId    string
	AirportIATA string
	Type        model.Nature
	From        time.Time
	To          time.Time
}

type SensorDataRepository interface {
	Store(sensor model.SensorData) error
}

type SensorDataMongoRepository struct {
	client *mongo.Client
}

func NewSensorDataRepository() SensorDataRepository {
	client := NewMongoClient()
	return &SensorDataMongoRepository{client: client}
}

func toBson(sensorData model.SensorData) bson.D {
	return bson.D{
		{"timestamp", sensorData.Timestamp},
		{"value", sensorData.Value},
		{"metadata", bson.D{
			{"sensorId", sensorData.SensorId},
			{"airportIATA", sensorData.AirportIATA},
			{"sensorType", sensorData.Nature},
		}},
	}
}

func (r *SensorDataMongoRepository) Store(sensor model.SensorData) error {
	coll := r.client.Database(DATABASE).Collection(COLLECTION)
	_, err := coll.InsertOne(context.Background(), toBson(sensor))
	if err != nil {
		return err
	}
	return nil
}
