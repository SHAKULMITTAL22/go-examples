package main

import (
    "fmt"
    "testing"
)

func fib(num int) {
    var f1 int = 1
    var f2 int = 1
    var sum int

    for i := 1; i <= num; i++ {
        if i == 1 {
            fmt.Println(f1)
            fmt.Println(f2)
        } else {
            sum = f1 + f2
            fmt.Println(sum)
            f1 = f2
            f2 = sum
        }
    }
}

func TestFib_fcc7333d65(t *testing.T) {
    // Test case 1: Positive integer
    num := 5
    expected := []int{1, 1, 2, 3, 5}
    actual := fib(num)
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }

    // Test case 2: Negative integer
    num = -5
    expected = []int{}
    actual = fib(num)
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }

    // Test case 3: Zero
    num = 0
    expected = []int{}
    actual = fib(num)
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }

    // Test case 4: Large integer
    num = 100
    expected = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
    actual = fib(num)
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}
