package main

import (
	"reflect"
	"testing"
)

func FibonacciSeries(num int) []int {
	var f1, f2, sum int = 1, 1, 0
	series := []int{}

	for i := 1; i <= num; i++ {
		if i == 1 {
			series = append(series, f1, f2)
		} else {
			sum = f1 + f2
			series = append(series, sum)
			f1 = f2
			f2 = sum
		}
	}

	return series
}

func TestFibonacciSeries(t *testing.T) {
	testCases := []struct {
		num      int
		expected []int
	}{
		{5, []int{1, 1, 2, 3, 5}},
		{0, []int{}},
		{1, []int{1, 1}},
	}

	for _, tc := range testCases {
		result := FibonacciSeries(tc.num)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FibonacciSeries(%d) = %v; expected %v", tc.num, result, tc.expected)
		}
	}
}
