# Airport Monitoring

## Database Recorder

### Description

Store every data received send by sensors on the mqtt broker in a mongo database.

### Run 

```bash
./database_recorder <path_to_yaml_config_file>
```

### Config file

```yaml
# Database recorder configuration file
datasource:
  url: mongodb://localhost:27017
  username:
  password:

mqtt:
  broker:
    url: tcp://localhost:1883
    username:
    password:
  client:
    id: database_recorder
```

### Environment variables

Instead of using a config file, you can use environment variables:
- DATASOURCE_URL
- DATASOURCE_USERNAME
- DATASOURCE_PASSWORD
- MQTT_BROKER_URL
- MQTT_BROKER_USERNAME
- MQTT_BROKER_PASSWORD
- MQTT_CLIENT_ID

## Add .env
WEATHER_URL="http://api.weatherstack.com/"
WEATHER_API_KEY=

### Measure Example : 
    - wind_speed
    - temperature
    - pressure
    - humidity
    - precip
    - visibility
    - is_day
    - uv_index

## To start new sensor with config file
cd cmd/sensor

go build -o name_measure/name_exe

cd name_measure

./name_exe --config=./config_file.json
