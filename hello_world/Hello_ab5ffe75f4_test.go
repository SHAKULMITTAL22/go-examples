package main

import (
	"testing"
	"bytes"
	"fmt"
	"os"
	"io"
)

func helloWorld() {
	fmt.Println("hello world!")
}

func TestHello_ab5ffe75f4(t *testing.T) {
	// Redirect stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	helloWorld()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old
	out := <-outC

	expected := "hello world!\n"
	if out != expected {
		t.Errorf("Expected %v, but got %v", expected, out)
	}
}
