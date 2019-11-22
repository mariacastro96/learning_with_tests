package shape

import "math"

// Rectangle has a width and a height values
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle has a radius value
type Circle struct {
	Radius float64
}

// Shape inplements area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Perimeter returns the parimeter of a shape
func (rectangle Rectangle) Perimeter() float64 {
	perimeter := 2 * (rectangle.Width + rectangle.Height)
	return perimeter
}

// Area returns the area of a rectangle
func (rectangle Rectangle) Area() float64 {
	area := rectangle.Width * rectangle.Height
	return area
}

// Area returns the area of a circle
func (circle Circle) Area() float64 {
	area := math.Pi * circle.Radius * circle.Radius

	return area
}
