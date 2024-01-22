export interface Sensor {
    sensorId: string,
    sensorType: string,
    avg?: number,
    airportIATA: string,
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