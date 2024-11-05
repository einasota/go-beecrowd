package main

import "fmt"

func main(){
	var A, B int
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	PROD := A * B
	fmt.Printf("PROD = %d\n", PROD)
}