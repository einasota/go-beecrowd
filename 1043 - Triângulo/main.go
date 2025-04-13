// Link: https://resources.beecrowd.com/repository/UOJ_1043.html
package main

import "fmt"

func main() {
	var a,b,c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	if a<b+c && b<a+c && c<a+b {
		fmt.Printf("Perimetro = %.1f\n", a + b + c)
	} else {
		fmt.Printf("Area = %.1f\n",	(a + b)/2 * c)
	}
}
