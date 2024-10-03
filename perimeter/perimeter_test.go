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

		assertEqual(t, got, want)
	})
	t.Run("cirles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 2 * math.Pi * 10

		assertEqual(t, got, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		assertEqual(t, got, want)
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100},
		{Circle{10}, 314.1592653589793},
	}

	for _, value := range areaTests {
		checkArea(t, value.shape, value.want)
	}
}

func assertEqual(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("Want %.2f, got %.2f", want, got)
	}
}
