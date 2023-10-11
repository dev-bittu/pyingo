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

func (s *Set) Add(element interface{}) {
  s.Items[element] = true
}

func (s *Set) Union(set2 Set) Set {
  final_set := Set{}
  for i, ok := range s.Items {
    if ok {
      final_set.Items[i] = true
    }
  }
  for i, ok := range set2.Items {
    if ok {
      final_set.Items[i] = true
    }
  }
  return final_set
}
