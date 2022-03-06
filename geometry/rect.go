package geometry

type rect struct {
	width, height float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}
func (r rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}
