package main

import (
	"testing"
	"bytes"
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func testFib(num int) {
	var f1 int = 1
	var f2 int = 1
	var sum int

	for i := 1; i <= num; i++ {
		if i == 1 {
			fmt.Fprintln(out, f1)
			fmt.Fprintln(out, f2)
		} else {
			sum = f1 + f2
			fmt.Fprintln(out, sum)
			f1 = f2
			f2 = sum
		}
	}
}

func TestFib_91dd7bc113(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	testFib(5)
	got := buf.String()
	want := "1\n1\n2\n3\n5\n"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFib_91dd7bc113_Zero(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	testFib(0)
	got := buf.String()
	want := ""

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
