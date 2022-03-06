package interfaces

type geometry interface {
	Area() float64
	Perimeter() float64
}

func MeasureArea(g geometry) (r float64) {
	return g.Area()
}

func MeasurePerimeter(g geometry) (r float64) {
	return g.Perimeter()
}
