package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool

const Pi = 3.14

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) float64 {
	return float64(x)*0.1 + 1.0
}

func main() {
	// basic function, specifying types
	x, y := 42, 13
	fmt.Printf("Watch me add %d + %d => %d\n", x, y, add(x, y))

	// multiple returned items
	start_a, start_b := "this", "that"
	end_a, end_b := swap(start_a, start_b)
	fmt.Printf("swap(%q, %q) => %q and %q\n", start_a, start_b, end_a, end_b)

	// named return value -- naked return statement
	split_input := 17
	s1, s2 := split(split_input)
	fmt.Printf("split(%d) => %d and %d\n", split_input, s1, s2)

	// variables
	var i int
	fmt.Println(i, c, python, java)

	// types
	const c = "%T(%v)\n"
	fmt.Printf(c, ToBe, ToBe)
	fmt.Printf(c, MaxInt, MaxInt)
	fmt.Printf(c, z, z)

	// zero values
	var j int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", j, f, b, s)

	// type assignment + usage
	intVar := 42
	floatVar := 42.0
	fmt.Printf("intVar is of type %T\n", intVar)
	fmt.Printf("floatVar is of type %T\n", floatVar)
	fmt.Printf("needInt(%v(%T)) => %v\n", intVar, intVar, needInt(intVar))

	// constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
