package main

import "fmt"

func main () {
	var a int
	var b, c float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	salario := b * c
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", a, salario)
}