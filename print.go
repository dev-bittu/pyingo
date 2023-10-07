package pyingo

import "fmt"

func Print(elements ...printer) {
  for _, element := range elements {
    fmt.Print(element)
    fmt.Print(" ")
  }
}

func Println(elements ...printer) {
  for _, element := range elements {
    fmt.Print(element)
    fmt.Print(" ")
  }
  fmt.Println()
}
