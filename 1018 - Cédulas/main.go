package main

import "fmt"

func main(){
	notas := [7]int{100, 50, 20, 10, 5, 2, 1}
	var valor int 

	fmt.Scanf("%d", &valor)

	for _, nota := range notas {
		quantidadeNotas := valor / nota
		fmt.Printf("%d nota(s) de R$ %d,00\n", quantidadeNotas, nota)
		valor = valor % nota
	}
}