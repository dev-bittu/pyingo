package pyingo

type Tuple struct {
	Items []interface{}
}

func (l *Tuple) String() string {
	repr := ""
	for _, i := range l.Items {
		repr += i.(string)
	}
	return "(" + repr + ")"
}

func (l *Tuple) Clear() {
	clear(l.Items)
}

func (l *Tuple) Count(element interface{}) int {
	occurance := 0
	for _, i := range l.Items {
		if i == element {
			occurance += 1
		}
	}
	return occurance
}

func (l *Tuple) Len() int {
	return len(l.Items)
}

func (l *Tuple) Index(element interface{}, start, end int) int {
	var index int
	if start == 0 && end == 0 {
		for i, j := range l.Items {
			if j == element {
				index = i
			}
		}
	} else {
		mylist := l.Items[start:end]
		for i, j := range mylist {
			if j == element {
				index = i
			}
		}
	}
	return index
}
