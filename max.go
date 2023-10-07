package pyingo

func Max(list List) interface{} {
  max := list.Items[0]
  for _, j := range list.Items {
    if j.(int) >= max.(int) {
      max = j
    }
  }
  return max
}
