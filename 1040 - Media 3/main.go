// Link: https://resources.beecrowd.com/repository/UOJ_1040.html
package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scanf("%f %f %f %f", &n1, &n2, &n3, &n4)
	media := (2*n1 + 3*n2 + 4*n3 + n4)/10
	fmt.Printf("Media: %.1f\n", media)
	if media >= 7.0 {
		fmt.Printf("Aluno aprovado.\n")
	} else if media < 5.0 {
		fmt.Printf("Aluno rrprovado.\n")
	} else {
		fmt.Printf("Aluno em exame.\n")
		var exame float64
		fmt.Scanf("%f", &exame)
		fmt.Printf("Nota do exame: %.1f\n", exame)
		media = (media + exame)/2
		if media >= 5.0{
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}
		fmt.Printf("Media final: %.1f\n", media)
	}
}
