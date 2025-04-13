// Link: https://resources.beecrowd.com/repository/UOJ_1038.html
package main

import "fmt"

func main() {
	var codigo, quantidade int
	fmt.Scanf("%d %d", &codigo, &quantidade)
	precos := [5]float64 {4.00, 4.50, 5.00, 2.00, 1.50}
	fmt.Printf("Total: R$ %.2f\n", float64(quantidade) * precos[codigo - 1])
}
