package main

import "fmt"

func ExampleAverage() {
  slice := []float64{1, 2, 3, 4, 5}

  var sum float64

  for _, i := range slice {
    sum += i
  }

  fmt.Println(sum / float64(len(slice)))

  // Output: 3

}
