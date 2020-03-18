package smi

import (
	"testing"
)

func assertResult(want float64, got float64, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("the perimeter of a rectangle given the height and width", func(t *testing.T) {
		r := Rectangle{23.7, 123.5}
		want := float64(294.4)
		got := Perimeter(r)
		assertResult(want, got, t)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{23.7, 123.5}, float64(2926.95)},
		{"circle", Circle{23.7}, float64(1764.829980)},
		{"triangle", Triangle{567, 889.1}, float64(252059.85)},
	}
	/* shapeArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	} */

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("[x] %v test case got %g want %g", test.name, got, test.want)
		}
	}

	/* t.Run("the area of a rectangle given the height and width", func(t *testing.T) {
		r := Rectangle{23.7, 123.5}
		shapeArea(t, r, float64(2926.95))
	})
	t.Run("the area of a circle given the radius", func(t *testing.T) {
		r := Circle{23.7}
		shapeArea(t, r, float64(1764.829980))
	}) */
}
