// Link => https://resources.beecrowd.com/repository/UOJ_1010.html
package main

import (
	"fmt"
)

func main() {
	var codigo1, qtd1, codigo2, qtd2 int
	var valor1, valor2 float64

	fmt.Scanf("%d %d %f", &codigo1, &qtd1, &valor1)
	fmt.Scanf("%d %d %f", &codigo2, &qtd2, &valor2)

	total := float64(qtd1)*valor1 + float64(qtd2)*valor2
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}