package pyingo

type Set struct {
  Items map[interface{}]bool
}

func (s *Set) String() string {
  repr := ""
  for i, ok := range s.Items {
    if ok {
      repr += i.(string)+" "
    }
  }
  return "{"+repr+"}"
}
