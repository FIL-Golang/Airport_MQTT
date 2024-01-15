package persist

import (
	"Airport_MQTT/internal/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "airport_sensors"
	COLLECTION = "sensor_data"
)

type SensorDataMongo struct {
	Timestamp primitive.DateTime
	Value     float32
	Metadata  SensorDataMongoMetadata
}

type SensorDataMongoMetadata struct {
	SensorId    string
	AirportIATA string
	Nature      model.Nature
}

type Filter struct {
	SensorId    string
	AirportIATA string
	Type        model.Nature
	From        time.Time
	To          time.Time
}

type SensorDataRepository interface {
	Store(sensor model.SensorData) error
	FindAllReadingsByFilter(filter Filter) ([]model.Sensor, error)
}

type SensorDataMongoRepository struct {
	client *mongo.Client
}

func NewSensorDataRepository() SensorDataRepository {
	client := NewMongoClient()
	return &SensorDataMongoRepository{client: client}
}

func (r *SensorDataMongoRepository) Store(sensor model.SensorData) error {
	coll := r.client.Database(DATABASE).Collection(COLLECTION)
	_, err := coll.InsertOne(context.Background(), bson.D{
		{"timestamp", sensor.Timestamp},
		{"value", sensor.Value},
		{"metadata", bson.D{
			{"sensorId", sensor.SensorId},
			{"airportIATA", sensor.AirportIATA},
			{"sensorType", sensor.Nature},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *SensorDataMongoRepository) FindAllReadingsByFilter(filter Filter) ([]model.Sensor, error) {
	coll := r.client.Database(DATABASE).Collection(COLLECTION)

	//prepare request filter
	mongoFilter := map[string]interface{}{}
	if filter.SensorId != "" {
		mongoFilter["metadata.sensorId"] = filter.SensorId
	}
	if filter.AirportIATA != "" {
		mongoFilter["metadata.airportIATA"] = filter.AirportIATA
	}
	if filter.Type != 0 {
		mongoFilter["metadata.sensorType"] = filter.Type
	}
	if !filter.From.IsZero() && !filter.To.IsZero() {
		mongoFilter["timestamp"] = map[string]interface{}{
			"$gte": filter.From,
			"$lte": filter.To,
		}
	}

	//execute request
	cursor, err := coll.Find(context.Background(), mongoFilter)
	if err != nil {
		return nil, err
	}

	sensors := mapSensorDataMongo(cursor)

	return sensors, nil
}

func mapSensorDataMongo(sensorData *mongo.Cursor) []model.Sensor {
	var sensors []model.Sensor
	for sensorData.Next(context.Background()) {
		var sensorDataMongo SensorDataMongo
		err := sensorData.Decode(&sensorDataMongo)
		if err != nil {
			fmt.Println("Error while decoding sensor data")
			continue
		}

		sensor := model.Sensor{
			SensorId:    sensorDataMongo.Metadata.SensorId,
			AirportIATA: sensorDataMongo.Metadata.AirportIATA,
			Type:        sensorDataMongo.Metadata.Nature,
		}

		reading := model.Reading{
			Timestamp: sensorDataMongo.Timestamp.Time(),
			Value:     sensorDataMongo.Value,
		}

		//check if sensor already exists
		exists := false
		for i, s := range sensors {
			if s.SensorId == sensor.SensorId {
				sensors[i].Readings = append(sensors[i].Readings, reading)
				exists = true
				break
			}
		}

		if !exists {
			sensor.Readings = []model.Reading{reading}
			sensors = append(sensors, sensor)
		}
	}
	return sensors
}
