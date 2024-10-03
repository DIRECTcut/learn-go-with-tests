package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{15.0, 15.0})
	want := 60.0

	assertEqual(got, want, t)
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{10.0, 10.0})
	want := 100.0

	assertEqual(got, want, t)
}

func assertEqual(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("Want %.2f, got %.2f", want, got)
	}
}
