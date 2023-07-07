package main

import "testing"

func TestIs_even_6dea0b9d5e(t *testing.T) {
    // Test case for even number
    num := 4
    expectedResult := 1
    if result := is_even(num); result != expectedResult {
        t.Errorf("Expected %d but got %d", expectedResult, result)
    }

    // Test case for odd number
    num = 5
    expectedResult = 0
    if result := is_even(num); result != expectedResult {
        t.Errorf("Expected %d but got %d", expectedResult, result)
    }

    // Test case for zero
    num = 0
    expectedResult = 1
    if result := is_even(num); result != expectedResult {
        t.Errorf("Expected %d but got %d", expectedResult, result)
    }

    // Test case for negative even number
    num = -6
    expectedResult = 1
    if result := is_even(num); result != expectedResult {
        t.Errorf("Expected %d but got %d", expectedResult, result)
    }

    // Test case for negative odd number
    num = -7
    expectedResult = 0
    if result := is_even(num); result != expectedResult {
        t.Errorf("Expected %d but got %d", expectedResult, result)
    }
}

func is_even(num int) int {
    if num%2 == 0 {
        return 1
    } else {
        return 0
    }
}
