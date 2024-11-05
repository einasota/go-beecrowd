package main

import "fmt"

func main() {
	var name string
	var salary, sells float64
	fmt.Scanln(&name)
	fmt.Scanln(&salary)
	fmt.Scanln(&sells)
	comission := sells * 0.15
	total := salary + comission
	fmt.Printf("TOTAL = R$ %.2f\n", total)
}