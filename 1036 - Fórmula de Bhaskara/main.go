// Link: https://resources.beecrowd.com/repository/UOJ_1036.html
package main

import (
  "fmt"
  "math"
)

func main() {
  var a, b, c float64
  fmt.Scanf("%f %f %f", &a, &b, &c)
  delta := (b*b)-4*a*c
	if a != 0 && delta > -1 {
		x1 := (-b + math.Sqrt(delta))/(2*a)
  	x2 := (-b-(math.Sqrt(delta)))/(2*a)
  	fmt.Printf("R1 = %.5f\nR2 = %.5f\n", x1, x2)
	} else {
		fmt.Printf("Impossivel calcular\n")
	}
}
