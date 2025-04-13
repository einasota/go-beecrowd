// Link: https://resources.beecrowd.com/repository/UOJ_1042.html
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int
	fmt.Scanf("%d %d %d", &n1, &n2, &n3)
	init := []int{n1, n2, n3}
	sort.Ints(init)
	fmt.Printf("%d\n%d\n%d\n", init[0], init[1], init[2])
	fmt.Printf("\n%d\n%d\n%d\n",n1,n2,n3)
}
