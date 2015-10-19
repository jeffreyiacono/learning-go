package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
	//"github.com/jeffreyiacono/stringutil"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// can use v here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}

func newtonSqrt(z, x float64) float64 {
	return z - ((math.Pow(z, 2) - x) / (2.0 * z))
}

func Sqrt(x float64) float64 {
	const ChangeThreshold float64 = 0.0001
	var z float64

	for i := 1; i <= 10; i++ {
		z = float64(i)
		for j := 0; j < 10; j++ {
			if new_z := newtonSqrt(z, float64(x)); math.Abs(new_z-z) <= ChangeThreshold {
				return new_z
			} else {
				fmt.Printf("starting z=%v, iteration=%v, z=%v, threshold (%v) not yet hit, continuing ...\n", i, j, z, ChangeThreshold)
				z = new_z
			}
		}
	}
	return z
}

func main() {
	defer fmt.Println("world")
	// loop basic
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// loops without starting / iterator conditions => while loop!
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// conditionals
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// loops + functions
	fmt.Printf("Newton: %v vs math.Sqrt: %v\n", Sqrt(2), math.Sqrt(2))

	// switch => os
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows ...
		fmt.Printf("%s.\n", os)
	}

	// switch => time
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch => hours
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// deferring ... see main top for defer statement
	fmt.Println("hello")
}
