import {describe, expect, test} from '@jest/globals';
import {hasIncreased, composeIncreasedDepths} from './solution';

describe('hasIncreased', () => {
    test('2 gt 1', () => {
        expect(hasIncreased(4,2)).toBe(true);
    });
});

describe('composeIncreasedDepths', () => {
    test('will count how many times depth increases', () => {
        expect(composeIncreasedDepths([1,2,3,3,3,4,2,1,5,2])).toBe(4);
    });
});


