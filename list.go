package pyingo

type List struct {
	Items []interface{}
}

func (l *List) String() string {
	repr := ""
	for _, i := range l.Items {
		repr += i.(string)
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
