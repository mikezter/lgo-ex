package main

import "fmt"

func fibo(n int) int {
  switch n {
  case 1, 2:
    fmt.Print("1 ")
    return 1
  default:
    v := fibo(n-1) + fibo(n-2)
    fmt.Printf("%v ", v)
    return v
  }
}

func ExampleFibo() {

  n := 30
  fibo(n)
  for i := n + 1; i < n+3; i++ {

  }
  // Output: 1 1 2 3 5 8 13
}
