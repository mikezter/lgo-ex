package main

import "fmt"

func collect(valid func(int) bool, list []int) []int {
  res := []int{}
  for _, v := range list {
    if valid(v) {
      res = append(res, v)
    }
  }
  return res
}

func ExampleMap() {
  i := []int{4, 2, 3, 4, 5, 5, 7, 8, 9, 10, 22}

  valid := func(i int) bool {
    return i%2 == 0
  }

  fmt.Println(collect(valid, i))

  // Output: [4 2 4 8 10 22]
}
