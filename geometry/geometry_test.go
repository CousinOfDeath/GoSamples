package geometry

import (
	"math"
	"math/interfaces"
	"testing"
)

func TestMeasureRectanle(t *testing.T) {
	r := rect{width: 3, height: 4}
	var result = interfaces.MeasureArea(r)
	expected := 12.0

	//c := circle{radius: 5}

	if result != expected {
		t.Errorf("Result %f and expected %f", result, expected)
	}
}

func TestPerimeterRectanle(t *testing.T) {
	r := rect{width: 3, height: 4}
	var result = interfaces.MeasurePerimeter(r)
	expected := 14.0

	if result != expected {
		t.Errorf("Result %f and expected %f", result, expected)
	}
}

func TestMeasureCircle(t *testing.T) {
	c := circle{radius: 5}
	var result = interfaces.MeasureArea(c)
	//Absolute comparison of floating points won't work due to rounding errors
	result = math.Round((result * 100) / 100)
	expected := 79.000000

	if result != expected {
		t.Errorf("Result %f and expected %f", result, expected)
	}
}

func TestPerimeterCircle(t *testing.T) {
	c := circle{radius: 5}
	var result = interfaces.MeasurePerimeter(c)
	//Absolute comparison of floating points won't work due to rounding errors
	result = math.Round((result * 100) / 100)
	expected := 31.000000

	if result != expected {
		t.Errorf("Result %f and expected %f", result, expected)
	}
}
