package ex03

import "testing"

func Test_SumOfSquaresLargestTwo(t *testing.T) {
	type testDataType struct {
		a, b, c float64
		result  float64
	}
	testData := []testDataType{
		{1, 2, 3, 13},
		{1, 3, 2, 13},
		{2, 1, 3, 13},
		{2, 3, 1, 13},
		{3, 1, 2, 13},
		{3, 2, 1, 13},
		{1, 1, 2, 5},
		{2, 1, 1, 5},
		{1, 2, 1, 5},
		{1, 1, 1, 2},
		{0, 0, 0, 0},
		{-1, -1, -1, 2},
		{-1, -2, -3, 5},
		{-1, -3, -2, 5},
		{-2, -1, -3, 5},
		{-2, -3, -1, 5},
		{-3, -1, -2, 5},
		{-3, -2, -1, 5},
		{-1, 2, 3, 13},
		{-1, 0, 1, 1},
		{-1, 1, 0, 1},
		{0, 1, -1, 1},
		{0, -1, 1, 1},
		{1, 0, -1, 1},
		{1, -1, 0, 1},
		{0, 0, 0.5, 0.25},
		{0, 0.5, 0.5, 0.5}}
	for _, data := range testData {
		result := SumOfSquaresLargestTwo(data.a, data.b, data.c)
		if result != data.result {
			t.Errorf("SumOfSquaresLargestTwo(%v, %v, %v) is %v, not %v\n",
				data.a, data.b, data.c, result, data.result)
		}
	}
}
