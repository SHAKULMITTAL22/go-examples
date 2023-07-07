package main

import (
	"testing"
	"os"
	"fmt"
	"bytes"
	"io"
)

func helloWorldTest() {
	fmt.Println("hello world!")
}

func TestMain(t *testing.T) {
	// Keep backup of the real stdout
	old := os.Stdout 
	r, w, _ := os.Pipe()
	os.Stdout = w

	helloWorldTest()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	expected := "hello world!\n"
	if out != expected {
		t.Errorf("Expected %v, but got %v", expected, out)
	}
}
