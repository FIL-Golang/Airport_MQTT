
# Airport Monitoring ✈️

## Architecture 

![Architecture Airport Monitoring](https://i.imgur.com/HrZwsRp.png)

## API REST

### Description
L'API REST offre une interface pour la gestion et l'interrogation des données. Elle permet aux utilisateurs et aux systèmes externes d'accéder aux données des capteurs et de visualiser les statistiques.

### Run 

```bash
./web_api <path_to_yaml_config_file>
```

## Frontend SPA 

### Description
Le Frontend SPA (Single Page Application) propose une interface utilisateur interactive pour visualiser les données en temps réel. Il comprend des graphiques dynamiques et des notifications pour les alertes.

### Run

```javascript
cd frontend-spa
npm run dev
```

## Database Recorder

### Description
Le Database Recorder enregistre toutes les données reçues des capteurs via le broker MQTT dans une base de données MongoDB. Cette composante assure la collecte, le stockage sécurisé et l'organisation des données pour des analyses futures.

### Run 

```bash
./database_recorder <path_to_yaml_config_file>
```

## File Recorder

### Description
Le File Recorder sauvegarde les données des capteurs dans des fichiers CSV. Cet outil est utile pour créer des archives de données ou pour des opérations où la base de données n'est pas disponible.

### Run 

```bash
./file_recorder <path_to_yaml_config_file>
```

## Alert Manager

### Description
L'Alert Manager gère les conditions d'alerte pour les données de capteurs. Il surveille les données en temps réel et déclenche des notifications en cas de dépassement de seuils prédéfinis.

### Run 

```bash
./alert_manager <path_to_yaml_config_file>
```

## Sensor

### Description
Le composant Sensor simule de la données qui provient d'une API. Envoi de données données tel que la vitesse du vent, la pression, la température, etc., au broker MQTT.

### Run 

```bash
./sensor <path_to_yaml_config_file>
```

## Config file

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
    qos: 0

sensor:
  airportIATA: 
  deviceId: 123e4567-e89b-12d3-a456-426655440000
  sensorType: <choose one on measure example>
  frequency: 30

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

### Measure Example: 
- wind_speed
- temperature
- pressure
- humidity
- precip
- visibility
- is_day
- uv_index

# Contributors 

- Titouan Cocheril
- Clément Repel
- Elias Morio
- Clément Galiot
- Guillaume Acamer
