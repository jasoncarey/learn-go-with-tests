package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * 100.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea

			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}
}

func CheckArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}

type Shape interface {
	Area() float64
}
