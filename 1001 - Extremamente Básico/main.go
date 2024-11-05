package main

import "fmt"

func main () {
	var A, B int
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	x := A + B
	fmt.Println("X =",x)
}