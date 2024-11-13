// Link => https://resources.beecrowd.com/repository/UOJ_1002.html
package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 = 3.14159
	var raio float64
	fmt.Scanln(&raio)
	area := n * (math.Pow(raio, 2))
	fmt.Printf("A=%.4f\n",area)
}