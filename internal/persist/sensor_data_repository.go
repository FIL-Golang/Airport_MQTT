package persist

import (
	"Airport_MQTT/internal/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "airport_sensors"
	COLLECTION = "sensor_data"
)

type SensorDataRepository interface {
	Store(sensor model.SensorData) error
	GetDistinctAirportCodes() ([]string, error)
	FindAllSensor(filter Filter) ([]model.Sensor, error)
	FindAllReading(filter Filter) ([]model.Sensor, error)
	GetAvg(filter Filter) ([]model.Average, error)
}

type SensorDataMongoRepository struct {
	client *mongo.Client
}

func NewSensorDataRepository() SensorDataRepository {
	client := NewMongoClient()
	return &SensorDataMongoRepository{client: client}
}

func toBson(sensorData model.SensorData) bson.M {
	return bson.M{
		"metadata": bson.M{
			"sensorId":    sensorData.SensorId,
			"airportIATA": sensorData.AirportIATA,
			"sensorType":  sensorData.Nature,
		},
		"value":     sensorData.Value,
		"timestamp": sensorData.Timestamp,
	}
}

func (r *SensorDataMongoRepository) Store(sensor model.SensorData) error {
	coll := r.getCollection()
	_, err := coll.InsertOne(context.Background(), toBson(sensor))
	if err != nil {
		return err
	}
	return nil
}

type Filter struct {
	SensorId    string
	AirportIATA string
	Type        model.Nature
	From        time.Time
	To          time.Time
}

func (r *SensorDataMongoRepository) GetDistinctAirportCodes() ([]string, error) {
	coll := r.getCollection()

	pipeline := mongo.Pipeline{
		bson.D{{"$group", bson.D{{"_id", "$metadata.airportIATA"}}}},
	}

	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(cursor, context.Background())

	var results []string
	for cursor.Next(context.Background()) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *SensorDataMongoRepository) FindAllSensor(filter Filter) ([]model.Sensor, error) {
	coll := r.getCollection()
	request := []bson.M{
		{"$match": buildFilter(filter)},
		{
			"$group": bson.M{
				"_id": bson.M{
					"sensorId":    "$metadata.sensorId",
					"airportIATA": "$metadata.airportIATA",
					"sensorType":  "$metadata.sensorType",
				}},
		},
		{
			"$project": bson.M{
				"_id":         0,
				"sensorId":    "$_id.sensorId",
				"airportIATA": "$_id.airportIATA",
				"sensorType":  "$_id.sensorType",
			},
		},
	}

	cursor, err := coll.Aggregate(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var res []model.Sensor
	err = cursor.All(context.Background(), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *SensorDataMongoRepository) FindAllReading(filter Filter) ([]model.Sensor, error) {
	coll := r.getCollection()
	request := []bson.M{
		{"$match": buildFilter(filter)},
		{
			"$group": bson.M{
				"_id": "$metadata",
				"readings": bson.M{
					"$push": bson.M{
						"timestamp": "$timestamp",
						"value":     "$value",
					},
				},
			},
		},
		{
			"$project": bson.M{
				"_id":         0,
				"sensorId":    "$_id.sensorId",
				"airportIATA": "$_id.airportIATA",
				"sensorType":  "$_id.sensorType",
				"readings":    1,
			},
		},
	}

	cursor, err := coll.Aggregate(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var res []model.Sensor
	err = cursor.All(context.Background(), &res)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *SensorDataMongoRepository) GetAvg(filter Filter) ([]model.Average, error) {
	coll := r.getCollection()
	request := []bson.M{
		{"$match": buildFilter(filter)},
		{
			"$group": bson.M{
				"_id": "$metadata.sensorType",
				"avg": bson.M{"$avg": "$value"},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"sensorType": "$_id",
				"avg":        1,
			},
		},
	}

	cursor, err := coll.Aggregate(context.Background(), request)
	fmt.Println(cursor)
	if err != nil {
		return nil, err
	}

	var res []model.Average
	err = cursor.All(context.Background(), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *SensorDataMongoRepository) getCollection() *mongo.Collection {
	return r.client.Database(DATABASE).Collection(COLLECTION)
}

// buildFilter builds a bson.D object from a Filter object
// It does not handle the case where the filter is empty
func buildFilter(filter Filter) bson.M {
	f := bson.M{}
	if filter.SensorId != "" {
		f["metadata.sensorId"] = filter.SensorId
	}
	if filter.AirportIATA != "" {
		f["metadata.airportIATA"] = filter.AirportIATA
	}
	if filter.Type != model.Undefined {
		f["metadata.sensorType"] = filter.Type
	}
	if (!filter.From.IsZero()) && (!filter.To.IsZero()) {
		f["timestamp"] = bson.M{
			"$gte": filter.From,
			"$lte": filter.To,
		}
	} else if !filter.From.IsZero() {
		f["timestamp"] = bson.M{
			"$gte": filter.From,
		}
	} else if !filter.To.IsZero() {
		f["timestamp"] = bson.M{
			"$lte": filter.To,
		}
	}
	return f
}
