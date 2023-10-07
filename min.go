package pyingo

func Min(list List) interface{} {
  min := list.Items[0]
  for _, j := range list.Items {
    if j.(int) <= min.(int) {
      min = j
    }
  }
  return min
}
