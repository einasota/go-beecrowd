// Link => https://resources.beecrowd.com/repository/UOJ_1930.html
package main

import "fmt"

func main(){
	var t1, t2, t3, t4 int
	fmt.Scanf("%d %d %d %d", &t1, &t2, &t3, &t4)
	total := (t1 + t2 + t3 + t4) - 3
	fmt.Printf("%d\n", total)
}