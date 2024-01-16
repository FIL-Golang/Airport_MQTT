import { useEffect } from 'react';
import mqtt, { MqttClient, IClientOptions } from 'mqtt';
import {toast} from "sonner";

interface MqttComponentProps {
    airport: string;
    sensorList: string[];
}

function MqttComponent(props: MqttComponentProps) {
    const url = process.env.REACT_APP_MQTT_URL;
    const username = process.env.REACT_APP_MQTT_USERNAME;
    const password = process.env.REACT_APP_MQTT_PASSWORD;

    if(url && username && password) {
        useEffect(() => {
            const clientId = 'mqtt-client-' + Math.random().toString(16);

            const options: IClientOptions = {
                keepalive: 60,
                clientId: clientId,
                username: username,
                password: password,
                protocolId: 'MQTT',
                protocolVersion: 5,
                clean: true,
                reconnectPeriod: 1000,
                connectTimeout: 30 * 1000,
            }
            const client: MqttClient = mqtt.connect(url, options);

            client.on('connect', () => {
                console.log('MQTT client connected');
                for (let sensor of props.sensorList) {
                    client.subscribe("airport/" + props.airport + "/sensor/" +sensor + '/#');
                }
            });

            client.on('message', (_topic: string, message: Buffer) => {
                toast("Nouvelle alerte ", {
                    description: 'Message : ' + message.toString(),
                    action: {
                        label: "Ok",
                        onClick: () => console.log("Ok"),
                    },
                })
            });

            return () => {
                client.end();
            };
        }, []);
    }
    else{
        console.log("Missing environment variables for MQTT connection");
    }

    return (
        <div className="Mqtt"></div>
    );
}

export default MqttComponent;
