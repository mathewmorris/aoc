import { measurements } from './input';

export const hasIncreased = (num1: number, num2: number): boolean => num1 > num2;

export const composeIncreasedDepths = (measurements: number[]): number => {
    return measurements.reduce((increased, measurement, idx, measurements) => {

        const prevMeasurement = measurements[idx - 1];

        if (!prevMeasurement) return increased;

        if (hasIncreased(measurement, prevMeasurement)) {
            return increased + 1;
        }

        return increased;
    }
    , 0);
}

export const hasSetIncreased = (set: number[], prev: number) => {
    return set > prev;
}

