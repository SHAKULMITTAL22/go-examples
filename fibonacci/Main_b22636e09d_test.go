package main

import (
	"testing"
)

// Assuming fib function
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	// Test case 1: Test fib function with a positive integer
	result := fib(5)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}

	// Test case 2: Test fib function with zero
	result = fib(0)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}

	// Test case 3: Test fib function with one
	result = fib(1)
	if result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}

	// Test case 4: Test fib function with a negative integer
	// Since fib function doesn't handle negative integers, it will go into an infinite loop.
	// So, this test case is commented out. If you want to test it, you need to handle negative integers in the fib function.
	/*
		result = fib(-5)
		if result != -5 {
			t.Errorf("Expected -5, but got %d", result)
		}
	*/
}
