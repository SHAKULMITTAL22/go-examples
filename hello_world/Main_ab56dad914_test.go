package main

import (
	"testing"
)

func helloWorld() string {
	return "Hello, world!"
}

func TestHelloWorld(t *testing.T) {
	result := helloWorld()
	if result != "Hello, world!" {
		t.Errorf("helloWorld() = %s; want 'Hello, world!'", result)
	}
}

func TestHelloWorldEmpty(t *testing.T) {
	result := helloWorld()
	if result == "" {
		t.Error("helloWorld() = ''; want non-empty string")
	}
}
