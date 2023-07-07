package main

import (
	"fmt"
)

// CheckEvenOdd checks if a number is even or odd and returns a string message.
func CheckEvenOdd(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("%d is even!", num)
	} else {
		return fmt.Sprintf("%d is odd!", num)
	}
}

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fmt.Println(CheckEvenOdd(num))
}
