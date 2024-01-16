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
    url: tls://2e0cec621e954a97b4917df783269b7e.s2.eu.hivemq.cloud:8883
    username:
    password:
  client:
    id: database_recorder

sensor:
  airportIATA: 
  deviceId: 123e4567-e89b-12d3-a456-426655440000
  sensorType: 
  frequency: 

api:
  url: http://api.weatherstack.com/current?access_key=%s&query=%s
  secretKey: 
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

go build -o name_exe

./name_exe config_file.yaml
