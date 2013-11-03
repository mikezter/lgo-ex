package main

import "fmt"

func fibo(n int) int {
  switch n {
  case 1, 2:
    return 1
  default:
    return fibo(n-1) + fibo(n-2)
  }
}

func ExampleFibo() {

  n := 30
  for i := n + 1; i < n+3; i++ {

    fmt.Print(fibo(i))
    fmt.Print(" ")
  }
  // Output: 1 1 2 3 5 8 13
}
