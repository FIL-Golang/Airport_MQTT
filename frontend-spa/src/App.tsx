import React, { useState, useEffect } from 'react';
import {Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue} from "@/components/ui/select";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import MQTTComponent from "@/components/mqtt";
import { LineChart, Line, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer, XAxis } from 'recharts';
import {Sensor, ReadingsPerDate, ColorMap, TransformedDataItem} from "../utils/types";
import {capitalizeFirstLetter, generateUniqueKey, getRandomColor} from "../utils/utils";

function App() {
    const [airportList, setAirportList] = useState([]);
    const [selectedAirport, setSelectedAirport] = useState('');
    const [sensorList, setSensorList] = useState<Sensor[]>([]);
    const [transformedData, setTransformedData] = useState<TransformedDataItem[]>([]);
    const [colors, setColors] = useState<ColorMap>({});
    const [sensorsLoaded, setSensorsLoaded] = useState(false);

    useEffect(() => {
        loadAirports().then(() => console.log("Airports loaded"));
    }, []);

    useEffect(() => {
        if (selectedAirport) {
            loadSensors().then(() => {setSensorsLoaded(true); console.log("Sensors loaded")});
        }
    }, [selectedAirport]);

    useEffect(() => {
        const newTransformedData: any[] | ((prevState: never[]) => never[]) = [];
        const newColors = { ...colors };

        sensorList.forEach(sensor => {
            if (sensor.readings) {
                Object.entries(sensor.readings).forEach(([date, value]) => {
                    let dataPoint = newTransformedData.find(point => point.date === date);
                    if (!dataPoint) {
                        dataPoint = { date };
                        newTransformedData.push(dataPoint);
                    }
                    dataPoint[sensor.sensorType] = value;

                    if (!newColors[sensor.sensorType]) {
                        newColors[sensor.sensorType] = getRandomColor();
                    }
                });
            }
        });
        newTransformedData.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime());

        setTransformedData(newTransformedData);
        setColors(newColors);
    }, [sensorList, colors]);

    const loadAirports = async () => {
        try {
            const response = await fetch(`${process.env.REACT_APP_API_URL}/distinctAirportCodes`);
            const airports = await response.json();
            setAirportList(airports);
            setSelectedAirport(airports[0]);
        } catch (error) {
            console.error("Error fetching airport list", error);
        }
    };

    const loadSensors = async () => {
        try {
            const response = await fetch(`${process.env.REACT_APP_API_URL}/sensorsForAirport?airportIATA=${selectedAirport}`);
            let sensorsList = await response.json();

            let promises = sensorsList.map(async (sensor: { sensorType: string; airportIATA: any; sensorId: any; }) => {
                try {
                    const response = await fetch(`${process.env.REACT_APP_API_URL}/lastReadingForSensor?sensorId=${sensor.sensorId}&airportIATA=${sensor.airportIATA}`);
                    const data = await response.json();
                    let lastReading = data.Value ? parseFloat(data.Value).toFixed(1) : 'Aucune donnée';

                    const responseReadings = await fetch(`${process.env.REACT_APP_API_URL}/readingsForSensor?sensorId=${sensor.sensorId}&airportIATA=${sensor.airportIATA}`);
                    const dataReadings = await responseReadings.json();

                    let readingsPerDate: ReadingsPerDate = {};
                    dataReadings.forEach((data: { readings: any[]; }) => {
                        data.readings.forEach((reading) => {
                            readingsPerDate[reading.timestamp.split('T')[0]] = reading.value;
                        });
                    });

                    return { ...sensor, sensorType: sensor.sensorType, lastReading: lastReading, readings: readingsPerDate };
                } catch (error) {
                    console.error("Error fetching data for sensor", error);
                    return { ...sensor, sensorType: sensor.sensorType, lastReading: 'Aucune donnée', readings: {} };
                }
            });

            let updatedSensors = await Promise.all(promises);
            setSensorList(updatedSensors);
        } catch (error) {
            console.error("Error fetching sensor list", error);
        }
    };

    const handleSelectChange = (value: React.SetStateAction<string>) => {
        setSelectedAirport(value);
        setSensorsLoaded(false);
    };

    return (
        <>
            <div className="w-full p-4">
                <div className="flex flex-col gap-6">
                    <div>
                        <Select onValueChange={handleSelectChange} value={selectedAirport}>
                            <SelectTrigger className="w-[220px]">
                                <SelectValue placeholder="Choisissez un aéroport" />
                            </SelectTrigger>
                            <SelectContent defaultValue={"NTE"}>
                                <SelectGroup>
                                    {airportList.map((airport) => (
                                        <SelectItem key={airport} value={airport}>
                                            <SelectLabel>{airport}</SelectLabel>
                                        </SelectItem>
                                    ))}
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                    </div>
                    {sensorsLoaded ? (
                        <MQTTComponent airport={selectedAirport} sensorList={sensorList} />
                    ) : (
                        <p>Loading sensors...</p>
                    )}
                    <div className="flex flex-row w-full gap-6">
                        {sensorList.map((sensor) => (
                            <Card key={generateUniqueKey(sensor.sensorId)}>
                                <CardHeader>
                                    <CardTitle>{capitalizeFirstLetter(sensor.sensorType)}</CardTitle>
                                    <CardDescription>Moyenne du capteur de {sensor.sensorType} sur la piste</CardDescription>
                                </CardHeader>
                                <CardContent>
                                    <h1 className="font-bold text-3xl">{sensor.lastReading}</h1>
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                    <div className="w-full">
                        <Card className="w-full">
                            <CardHeader>
                                <CardTitle>Statistiques</CardTitle>
                                <CardDescription>
                                    Données en temps réel de tous les capteurs de l'aéroport
                                </CardDescription>
                            </CardHeader>
                            <CardContent>
                                <ResponsiveContainer width="100%" height={300}>
                                    <LineChart
                                        data={transformedData}
                                        margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
                                        <XAxis dataKey="date" />
                                        <CartesianGrid strokeDasharray="3 3" />
                                        <YAxis yAxisId="left" orientation="left" />
                                        <Tooltip />
                                        <Legend />
                                        {sensorList.map((sensor, index) => (
                                            <Line
                                                key={index}
                                                yAxisId="left"
                                                type="monotone"
                                                dataKey={sensor.sensorType}
                                                stroke={colors[sensor.sensorType]}
                                                activeDot={{ r: 8 }}
                                            />
                                        ))}
                                    </LineChart>
                                </ResponsiveContainer>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </div>
        </>
    );
}

export default App;
