package main

import "fmt"

func is_even_test(num int) int {
	if num%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func printEvenOrOddTest(num int) {
	i := is_even_test(num)

	if i == 1 {
		fmt.Println("number:", num, "is even")
	} else {
		fmt.Println("number:", num, "is odd")
	}
}

func main() {
	var num int

	fmt.Print("enter number:")
	fmt.Scan(&num)

	printEvenOrOddTest(num)
}
