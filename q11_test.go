package main

import "fmt"

func fibo(n int) int {
  switch n {
  case 1, 2:
    return 1
  default:
    v := fibo(n-1) + fibo(n-2)
    return v
  }
}

func ExampleFibo() {

  n := 1
  for i := n + 1; i < n+8; i++ {
    v := fibo(i)
    fmt.Printf("%v ", v)
  }
  // Output: 1 1 2 3 5 8 13
}
