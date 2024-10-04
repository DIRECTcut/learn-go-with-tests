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
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, hasArea: 100},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func assertEqual(t *testing.T, got float64, want float64) {
	if got != want {
		t.Errorf("Want %.2f, got %.2f", want, got)
	}
}
