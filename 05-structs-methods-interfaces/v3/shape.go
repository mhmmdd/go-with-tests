package shape

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Rectangle) Perimeter() float64 {
	return 2 * (c.Width + c.Height)
}

func (c Rectangle) Area() float64 {
	return c.Width * c.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
