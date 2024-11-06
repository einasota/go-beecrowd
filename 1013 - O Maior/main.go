package main

import (
	"fmt"
	"math"
)

func main()  {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	maiorab := (a+b+math.Abs(a-b))/2
	maiorc := (maiorab+c+math.Abs(maiorab-c)) / 2
	fmt.Printf("%.0f eh o maior\n", maiorc)
}