package geometry

import "math"

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
