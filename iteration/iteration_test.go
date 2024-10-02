package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("expected %s but got %s", got, want)
	}
}

func ExampleRepeat() {
	res := Repeat("b", 3)
	fmt.Println(res)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
