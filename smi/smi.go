package smi

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (c Circle) Area() float64 {
	return float64(3.142 * c.Radius * c.Radius)
}

func (t Triangle) Area() float64 {
	return float64(0.5 * t.Base * t.Height)
}

//Perimeter returns the perimeter size of a rectangle
func Perimeter(r Rectangle) float64 {
	return float64(2 * (r.Width + r.Height))
}

//Area returns the area of a rectangle
/* func Area(r Rectangle) float64 {
	return float64(r.Width * r.Height)
} */
