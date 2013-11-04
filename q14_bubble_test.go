package main

import "fmt"
import "testing"

var i = []int{99, 4, 2, 3, 4, 5, 5, 7, 8, 9, 10, 22}

func swap(list []int, a, b int) {
  var i, j = list[a], list[b]
  list[b], list[a] = i, j
}

func bubbleSort(list []int) []int {
  n := len(list)
  if n < 2 {
    return list
  }

  for i := 0; i < n-1; i++ {
    for j := i + 1; j < n; j++ {

      if list[j] > list[i] {
        list[i], list[j] = list[j], list[i]
      }

    }
  }

  return list
}

func ExampleBuble() {

  fmt.Println(bubbleSort(i))

  // Output: [99 22 10 9 8 7 5 5 4 4 3 2]
}

func TestSwap(t *testing.T) {
  a := []int{2, 10}
  swap(a, 0, 1)

  if !(a[0] == 10 && a[1] == 2) {
    t.Fatal("isnt swapped")
  }
}
