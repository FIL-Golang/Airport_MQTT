import useMQTT from '../hook/useMqtt.tsx';

const MQTTComponent = () => {
    const url = process.env.REACT_APP_MQTT_URL;
    const clientId = process.env.REACT_APP_MQTT_CLIENT_ID;
    const topic = process.env.REACT_APP_MQTT_TOPIC;

    if(url && clientId && topic) {
        useMQTT(url, clientId, topic);
    }

    return <div>Connexion MQTT en cours...</div>;
};

export default MQTTComponent;
