// Link => https://resources.beecrowd.com/repository/UOJ_1005.html
package main

import "fmt"

func main() {
	var A, B float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	media := ((A * 3.5) + (B * 7.5)) / 11.0
	fmt.Printf("MEDIA = %.5f\n", media)
}