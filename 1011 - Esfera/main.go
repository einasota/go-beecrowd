package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	var pi float64 = 3.14159
	fmt.Scanf("%f", &r)
	calculo := (4/3.0) * pi * math.Pow(r,3)
	fmt.Printf("VOLUME = %.3f\n", calculo)
}