package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	soma := A + B
	fmt.Printf("SOMA = %d\n", soma)
}