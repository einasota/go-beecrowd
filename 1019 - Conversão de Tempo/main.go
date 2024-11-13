// Link => https://resources.beecrowd.com/repository/UOJ_1019.html
package main

import "fmt"

func main() {
	var valor int
	fmt.Scanf("%d", &valor)
	horas := valor / 3600
	valor -= horas * 3600
	minutos := valor / 60
	valor -= minutos * 60
	segundos := valor / 1

	fmt.Printf("%d:%d:%d\n",horas,minutos,segundos)
}