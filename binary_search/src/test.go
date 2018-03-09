package main

import (
  "fmt"
  "bina"
)

func main() {
  list := []int{1, 3, 4, 5, 6, 7, 17, 20, 21, 22, 30}
  fmt.Println(list)
  fmt.Println("Is 6 in the list?", bina.BinarySearch(list, 6))
  fmt.Println("Is 18 in the list?", bina.BinarySearch(list, 18))
}
