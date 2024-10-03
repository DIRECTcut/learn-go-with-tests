package perimeter

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{15.0, 15.0}
		got := rectangle.Perimeter()
		want := 60.0

		assertEqual(got, want, t)
	})
	t.Run("cirles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 2 * math.Pi * 10

		assertEqual(got, want, t)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		assertEqual(got, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := math.Pi * math.Pow(10, 2)

		assertEqual(got, want, t)
	})
}

func assertEqual(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("Want %.2f, got %.2f", want, got)
	}
}
