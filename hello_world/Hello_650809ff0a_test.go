package main

import (
    "fmt"
    "testing"
)

func hello() {
    fmt.Println("hello world!")
}

func TestHello(t *testing.T) {
    // Test case 1: Check if the function prints the correct message.
    expected := "hello world!"
    actual := hello()
    if actual != expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }

    // Test case 2: Check if the function panics.
    func() {
        defer func() {
            if r := recover(); r == nil {
                t.Errorf("Expected panic, but none occurred")
            }
        }()
        hello()
    }()
}
