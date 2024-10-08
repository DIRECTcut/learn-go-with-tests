package countdown

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	fmt.Printf("First: %p\n", buffer)
	fmt.Printf("Second: %p\n", &buffer)

	Countdown(buffer)

	got := buffer.String()
	want := "3"
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
