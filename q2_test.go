package main

import (
  "fmt"
)

func ExampleArray() {

  a := make([]int, 10)
  for i := 0; i < 10; i++ {
    a[i] = i
  }
  fmt.Printf("%v ", a)
  // Output: foobar[0 1 2 3 4 5 6 7 8 9]
}

func ExampleForloop() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%v ", i)
  }
  // Output: 0 1 2 3 4 5 6 7 8 9
}

func ExampleGoto() {
  i := 0
Loop:
  fmt.Printf("%v ", i)
  i++
  if i < 10 {
    goto Loop
  }
  // Output: 0 1 2 3 4 5 6 7 8 9
}
