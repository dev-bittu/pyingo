package pyingo

type List struct {
  list []interface{}
}

func (l *List) String() string {
  repr := ""
  for _, i := range l.list {
    repr += i.(string)
  }
  return repr
}

func (l *List) Append(element interface{}) {
  l.list = append(l.list, element)
}

func (l *List) Pop() interface{} {
  if len(l.list) == 0 {
    panic("IndexError: list is empty")
  }
  index := len(l.list)-1
  element := l.list[index]
  l.list = l.list[:index]
  return element
}

func (l *List) Clear() {
  clear(l.list)
}
