package math

import (
	"fmt"
	"testing"
)

type addTest struct {
	num1, num2, expected int
}

var addTestCases = []addTest{
	{1, 1, 2},
	{2, 2, 4},
	{100, 1, 101},
	{99, 1, 100},
	{30, 45, 75},
}

var subtractTestCases = []addTest{
	{1, 1, 0},
	{2, 1, 1},
	{100, 1, 99},
	{99, 16, 83},
	{30, 45, -15},
}

func TestAdd(t *testing.T) {

	for _, test := range addTestCases {
		if result := Add(test.num1, test.num2); result != test.expected {
			t.Errorf("Result %d and expected %d", result, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {

	for _, test := range subtractTestCases {
		if result := subtract(test.num1, test.num2); result != test.expected {
			t.Errorf("Result %d and expected %d", result, test.expected)
		}
	}
}

//Benchmark test
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

// Go code docuumentation with examples
func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
