export interface Sensor {
    sensorId: string,
    sensorType: string,
    airportIATA: string,
    lastReading?: number,
    readings?: string[],
}

export type ReadingsPerDate = {
    [date: string]: number;
};

export interface ColorMap {
    [key: string]: string;
}

export interface TransformedDataItem {
    date: string;
    temperature: number;
}