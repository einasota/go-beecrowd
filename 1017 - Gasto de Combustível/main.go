// Link => https://resources.beecrowd.com/repository/UOJ_1017.html
package main

import "fmt"

func main() {
	var tempo, velocidade float64
    var gasto float64 = 12

    fmt.Scanf("%f",&tempo)
    fmt.Scanf("%f",&velocidade)

    total := (tempo * velocidade) / gasto

    fmt.Printf("%.3f\n",total)
}