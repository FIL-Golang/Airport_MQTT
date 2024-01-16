import { useEffect } from 'react';
import { Client } from 'paho-mqtt';

const useMQTT = (url: string, clientId: string, topic: string) => {
    useEffect(() => {
        const client = new Client(url, clientId);

        const onConnect = () => {
            console.log("Connecté au broker MQTT");
            client.subscribe(topic);
        };

        const onMessageArrived = (message: { payloadString: any; }) => {
            console.log("Message reçu :", message.payloadString);
        };

        client.onConnectionLost = responseObject => console.log("Connexion perdue :", responseObject.errorMessage);
        client.onMessageArrived = onMessageArrived;

        client.connect({ onSuccess: onConnect });

        return () => {
            if (client.isConnected()) {
                client.disconnect();
            }
        };
    }, [url, clientId, topic]);
};

export default useMQTT;
