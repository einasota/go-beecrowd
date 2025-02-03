// Link: https://resources.beecrowd.com/repository/UOJ_1036.html
package main

import (
  "fmt"
  "math"
)

func main() {
  var a, b, c float64
  fmt.Scanf("%f %f %f", &a, &b, &c)
  delta := (math.Pow(b,2))-4*a*c
  x1 := (-b+(math.Sqrt(delta)))/2*a
  fmt.Println(x1)
  //if err != nil {
  //  fmt.Println("Impossivel calcular")
  //}
  x2 := (-b-(math.Sqrt(delta)))/2*a
  fmt.Printf("R1 = %.5f\nR2 = %.5f\n", x1, x2)
}
