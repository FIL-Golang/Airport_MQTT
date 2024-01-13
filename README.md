# Airport Monitoring

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
