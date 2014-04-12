package ex03_test

import (
	"sicp/ch01/ex03"
	"testing"
)

func Test3DistinctValues(t *testing.T) {
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
		{3, 2, 1, 13}}
	for _, data := range testData {
		if ex03.SumOfSquaresLargestTwo(data.a, data.b, data.c) != data.result {
			t.Errorf("SumOfSquaresLargestTwo(%v, %v, %v) != %v\n",
				data.a, data.b, data.c, data.result)
		}
	}
}
