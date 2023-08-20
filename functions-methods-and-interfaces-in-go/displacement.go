package main

import (
	"fmt"
)

var a, v0, s0, t float64

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
	return fn
}

func main() {
	fmt.Print("Please enter the value of the acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Please enter the value of the initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Please enter the value of the initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Please enter the value of the time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Displacement after", t, "seconds:", fn(t))
}
