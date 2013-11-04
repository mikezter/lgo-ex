package main

import (
  "fmt"
  "testing"
)

func average(slice []float64) float64 {
  if len(slice) == 0 {
    return 0.0
  }

  sum := 0.0
  for _, v := range slice {
    sum += v
  }

  return sum / float64(len(slice))
}

func TestAverage(t *testing.T) {
  mySlice := []float64{1, 2, 3, 4, 5}

  avg := average(mySlice)
  if avg != 3 {
    t.Fatal("Average should be 3, but is", average)
  }

}

func order(i, j int) (a, b int) {
  if i < j {
    a, b = i, j
  } else {
    a, b = j, i
  }
  return
}

func TestOrder(t *testing.T) {
  a, b := order(5, 6)
  if a > b {
    t.Fatal("Should be ordered", a, b)
  }
  a, b = order(6, 5)
  if a > b {
    t.Fatal("Should be ordered", a, b)
  }
}

type stack struct {
  data [10]byte
  i    int
}

func (s *stack) push(i byte) {
  s.data[s.i] = i
  s.i += 1
}

func (s *stack) pop() (v byte) {
  if s.i == 0 {
    return 0
  }

  v = s.data[s.i-1]
  s.i -= 1

  return
}

func (s stack) String() (res string) {
  for i := 0; i < s.i; i++ {
    res += fmt.Sprintf("[%d:%s] ", i, string(s.data[i]))
  }
  return
}

func TestStack(t *testing.T) {
  var s stack

  s.push('m')

  a := s.pop()

  if a != 'm' {
    t.Fatal("Should have popped 'm', got", a)
  }

  s.push('m')
  s.push('s')

  fmt.Printf("The Stack: %v\n", s)

}

func varargs(args ...string) {
  for _, arg := range args {
    fmt.Println(arg)
  }
}
func ExampleVarargs() {
  varargs("muh", "foo", "toto")
  // Output:
  // muh
  // foo
  // toto
}

