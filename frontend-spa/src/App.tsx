import './App.css'
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue
} from "@/components/ui/select.tsx";
import MQTTComponent from "@/components/mqtt.tsx";

import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import { LineChart, Line, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import {useState} from "react";


function App() {
    const [selectedAirport, setSelectedAirport] = useState("NTE");
    const handleSelectChange = (value: string) => {
        setSelectedAirport(value);
    };

    // TODO : CALL API, GET SENSORS LIST, GET DATA

    const sensorList = ['temperature','pressure','wind_speed'];

    const data = [
        {
            name: 'Page A',
            uv: 4000,
            pv: 2400,
            amt: 2400,
        },
        {
            name: 'Page B',
            uv: 3000,
            pv: 1398,
            amt: 2210,
        },
        {
            name: 'Page C',
            uv: 2000,
            pv: 9800,
            amt: 2290,
        },
        {
            name: 'Page D',
            uv: 2780,
            pv: 3908,
            amt: 2000,
        },
        {
            name: 'Page E',
            uv: 1890,
            pv: 4800,
            amt: 2181,
        },
        {
            name: 'Page F',
            uv: 2390,
            pv: 3800,
            amt: 2500,
        },
        {
            name: 'Page G',
            uv: 3490,
            pv: 4300,
            amt: 2100,
        },
    ];
  return (
      <><MQTTComponent airport={selectedAirport} sensorList={sensorList}/>
          <div className="w-full">
              <div className="flex flex-col gap-6">
                  <div>
                      <Select onValueChange={handleSelectChange} value={selectedAirport}>
                          <SelectTrigger className="w-[220px]">
                              <SelectValue placeholder="Choissisez un aéroport"/>
                          </SelectTrigger>
                          <SelectContent defaultValue={"NTE"}>
                              <SelectGroup>
                                  <SelectLabel>Aéroports</SelectLabel>
                                  <SelectItem value="NTE">Nantes</SelectItem>
                                  <SelectItem value="ORY">Paris Orly</SelectItem>
                                  <SelectItem value="CDG">Paris Charles de Gaulle</SelectItem>
                                  <SelectItem value="DBX">Dubai</SelectItem>
                              </SelectGroup>
                          </SelectContent>
                      </Select>
                  </div>
                  <div className="flex flex-row w-full gap-6">
                      <Card className="w-1/3">
                          <CardHeader>
                              <CardTitle>Température</CardTitle>
                              <CardDescription>Capteurs de température sur la piste</CardDescription>
                          </CardHeader>
                          <CardContent>
                              <h1 className="font-bold text-3xl">22°C</h1>
                          </CardContent>
                      </Card>
                      <Card className="w-1/3">
                          <CardHeader>
                              <CardTitle>Pression</CardTitle>
                              <CardDescription>Capteurs de pression sur la piste</CardDescription>
                          </CardHeader>
                          <CardContent>
                              <h1 className="font-bold text-3xl">12 Pa</h1>
                          </CardContent>
                      </Card>
                      <Card className="w-1/3">
                          <CardHeader>
                              <CardTitle>Vitesse du vent</CardTitle>
                              <CardDescription>Capteurs de vitesse du vent sur la piste</CardDescription>
                          </CardHeader>
                          <CardContent>
                              <h1 className="font-bold text-3xl">16 km/h</h1>
                          </CardContent>
                      </Card>
                  </div>
                  <div className="w-full">
                      <Card className="w-full">
                          <CardHeader>
                              <CardTitle>Statistiques</CardTitle>
                              <CardDescription>Données en temps réel de tous les capteurs de
                                  l'aéroport</CardDescription>
                          </CardHeader>
                          <CardContent>
                              <ResponsiveContainer width="100%" height={300}>
                                  <LineChart
                                      width={500}
                                      height={300}
                                      data={data}
                                      margin={{
                                          top: 5,
                                          right: 30,
                                          left: 20,
                                          bottom: 5,
                                      }}
                                  >
                                      <CartesianGrid strokeDasharray="3 3"/>
                                      <YAxis yAxisId="left" orientation="left" stroke="#8884d8"/>
                                      <YAxis yAxisId="right" orientation="right" stroke="#82ca9d"/>
                                      <YAxis yAxisId="right" orientation="right" stroke="#FFD500"/>
                                      <Tooltip/>
                                      <Legend/>
                                      <Line yAxisId="left" type="monotone" dataKey="pv" stroke="#8884d8"
                                            activeDot={{r: 8}}/>
                                      <Line yAxisId="right" type="monotone" dataKey="uv" stroke="#82ca9d"/>
                                      <Line yAxisId="right" type="monotone" dataKey="amt" stroke="#FFD500"/>
                                  </LineChart>
                              </ResponsiveContainer>
                          </CardContent>
                      </Card>
                  </div>
              </div>
          </div>
      </>
  )
}

export default App
