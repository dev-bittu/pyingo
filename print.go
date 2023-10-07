package pyingo

import "fmt"

func Print(elements ...stringer) {
  for _, element := range elements {
    fmt.Print(element)
    fmt.Print(" ")
  }
}

func Println(elements ...stringer) {
  for _, element := range elements {
    fmt.Print(element)
    fmt.Print(" ")
  }
  fmt.Println()
}
