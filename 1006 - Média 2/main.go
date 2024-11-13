// Link => https://resources.beecrowd.com/repository/UOJ_1006.html
package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	media := (((a*2) + (b*3) + (c*5))/10.0)
	fmt.Printf("MEDIA = %.1f\n", media)
}