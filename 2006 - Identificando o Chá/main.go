// Link: https://resources.beecrowd.com/repository/UOJ_2006.html
package main

import "fmt" 

func main() {
  var t int
  total := 0
  guest := [5]int{}
  fmt.Scanf("%d", &t)
  fmt.Scanf("%d %d %d %d %d", &guest[0], &guest[1], &guest[2], &guest[3], &guest[4])
  for i := 0; i < len(guest); i++ {
    if t == guest[i] {
      total++
    }
  }
  fmt.Println(total)
}

