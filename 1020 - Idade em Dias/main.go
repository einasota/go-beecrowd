// Link => https://resources.beecrowd.com/repository/UOJ_1020.html
package main

import "fmt"

func main() {
	var totalDias int
	fmt.Scanf("%d", &totalDias)
	anos := totalDias/365
	meses := (totalDias%365) / 30
	dias := (totalDias%365) % 30
	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", anos, meses, dias)
}