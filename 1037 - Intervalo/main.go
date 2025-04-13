// Link: https://resources.beecrowd.com/repository/UOJ_1037.html
package main

import "fmt"

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	if 0 <= input && input <= 25 {
		fmt.Printf("Intervalo [0,25]\n")
	} else if 25 < input && input <= 50 {
		fmt.Printf("Intervalo (25,50]\n")
	} else if 50 < input && input <= 75 {
		fmt.Printf("Intervalo (50,75]\n")
	} else if 75 < input && input <= 100 {
		fmt.Printf("Intervalo (75,100]\n")
	} else {
		fmt.Printf("Fora de intervalo\n")
	}
}
