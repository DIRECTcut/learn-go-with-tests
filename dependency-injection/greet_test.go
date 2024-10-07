package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Chris")

	got := "Hello, Chris"
	want := buffer.String()

	if got != want {
		t.Errorf("Want %q got %q", want, got)
	}
}
