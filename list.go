package pyingo

type List struct {
	Items []interface{}
}

func (l *List) String() string {
	repr := ""
	for _, i := range l.Items {
		repr += i.(string)+" "
	}
	return "[" + repr + "]"
}

func (l *List) Append(element interface{}) {
	l.Items = append(l.Items, element)
}

func (l *List) Pop() interface{} {
	if len(l.Items) == 0 {
		panic("IndexError: list is empty")
	}
	index := len(l.Items) - 1
	element := l.Items[index]
	l.Items = l.Items[:index]
	return element
}

func (l *List) Clear() {
	clear(l.Items)
}

func (l *List) Extend(another_list List) {
	l.Items = append(l.Items, another_list.Items...)
}

func (l *List) Count(element interface{}) int {
	occurance := 0
	for _, i := range l.Items {
		if i == element {
			occurance += 1
		}
	}
	return occurance
}

func (l *List) Len() int {
	return len(l.Items)
}

func (l *List) Index(element interface{}, start, end int) int {
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

func (l *List) Remove(element interface{}) {
	var mylist []interface{}
	for _, i := range l.Items {
		if i != element {
			mylist = append(mylist, i)
		}
	}
	l.Items = mylist
}
