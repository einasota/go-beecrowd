// Link: https://resources.beecrowd.com/repository/UOJ_1035.html
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	if b > c && d > a && c + d > a + b && c > 0 && d > 0 && a % 2 == 0 {
		fmt.Printf("Valores aceitos\n")
	} else {
		fmt.Printf("Valores nao aceitos\n")
	}
}
