package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func NewList() List {
	return &list{}
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{
		Next:  l.first,
		Value: v,
	}

	if l.first == nil {
		l.last = i
	} else {
		l.first.Prev = i
	}

	l.first = i
	l.len++
	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{
		Prev:  l.last,
		Value: v,
	}

	if l.last == nil {
		l.first = i
	} else {
		l.last.Next = i
	}

	l.last = i
	l.len++
	return i
}

func (l *list) Remove(i *ListItem) {
	if l.len < 2 {
		l.first = nil
		l.last = nil
		l.len = 0
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.first = i.Next
		i.Next.Prev = nil
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.last = i.Prev
		i.Prev.Next = nil
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
