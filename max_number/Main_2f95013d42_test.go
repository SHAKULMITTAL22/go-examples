package main

import (
	"testing"
)

func findMax(numbers []int) int {
	var max int = 0
	for _, num := range numbers {
		if max < num {
			max = num
		}
	}
	return max
}

func TestFindMax(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	max := findMax(numbers)
	if max != 5 {
		t.Errorf("Max was incorrect, got: %d, want: %d.", max, 5)
	}
}
