// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison






package main

import "fmt"

func Fibonacci(n int) []int {
	x, y := 0, 1

	fibs := make([]int, n+1) // We add one to the length so that we can access the nth element of the slice directly using its index.
	fibs[0] = x
	fibs[1] = y

	// Fill in the remaining elements of the slice using a loop.
	for i := 2; i < len(fibs); i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	return fibs
}

func main() {
	fmt.Println("Enter number of terms:")
	var n int
	_, err := fmt.Scan(&n)
	if err!= nil {
		panic(err)
	}

	fibs := Fibonacci(n)
	for i, v := range fibs {
		fmt.Printf("%dth term is %d\n", i, v)
	}
}