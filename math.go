package math

import "time"

func Add(x, y int) (res int) {
	time.Sleep(10 * time.Microsecond)
	return x + y
}

func subtract(x, y int) (res int) {
	return x - y
}
