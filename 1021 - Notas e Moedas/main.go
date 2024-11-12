package main

import (
	"fmt"
	"math"
)

func main(){
	notas := [6]float64{100, 50, 20, 10, 5, 2}
	moedas := [6]float64{1, 0.5, 0.25, 0.10, 0.05, 0.01}
	
	var valor float64
	fmt.Scanf("%f", &valor)
	fmt.Printf("NOTAS:\n")
	for _, nota := range notas {
		quantidadeNotas := math.Floor(valor/nota)
		fmt.Printf("%.0f nota(s) de R$ %.0f.00\n", quantidadeNotas, nota)
		valor = math.Mod(valor, nota)
	}
	fmt.Printf("MOEDAS:\n")
	for _, moeda := range moedas {
		quantidadeMoedas := int(valor/moeda)
		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantidadeMoedas, moeda)
		valor -= float64(quantidadeMoedas) * moeda
		valor = math.Round(valor*100)/100
		// fmt.Printf("%f, %f", valor, quantidadeMoedas)
	}
}