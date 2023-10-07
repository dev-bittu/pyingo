package pyingo

func Sum(items List) int {
	sum := 0
	for _, i := range items.Items {
		sum += i.(int)
	}
	return sum
}
