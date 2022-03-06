package math

import (
	printpkg "math/print"
	"os"
	"time"
)

func Add(x, y int) (res int) {
	time.Sleep(10 * time.Microsecond)
	printpkg.PrintLine("Foooz", os.Stdout)
	return x + y
}

func Subtract(x, y int) (res int) {
	printpkg.PrintLine("Barz", os.Stdout)
	return x - y
}
