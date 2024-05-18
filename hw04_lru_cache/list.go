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
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := ListItem{
		Value: v,
		Next:  l.head,
		Prev:  nil,
	}
	if l.head != nil {
		l.head.Prev = &newListItem
	}
	l.head = &newListItem

	if l.tail == nil {
		l.tail = &newListItem
	}
	l.len++
	return &newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.tail,
	}

	if l.tail != nil {
		l.tail.Next = &newListItem
	}
	l.tail = &newListItem

	if l.head == nil {
		l.head = &newListItem
	}
	l.len++
	return &newListItem
}

func (l *list) Remove(i *ListItem) {
	prevItem := i.Prev
	nextItem := i.Next
	if prevItem != nil {
		prevItem.Next = nextItem
	}
	if nextItem != nil {
		nextItem.Prev = prevItem
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	prevItem := i.Prev // nil
	nextItem := i.Next // 60
	if prevItem != nil {
		prevItem.Next = nextItem
	} else if nextItem != nil {
		l.head = nextItem
	}

	if nextItem != nil {
		nextItem.Prev = prevItem
	} else if prevItem != nil {
		l.tail = prevItem
	}
	i.Next = l.head
	i.Prev = nil

	l.head.Prev = i
	l.head = i
}

func NewList() List {
	return &list{
		len:  0,
		head: nil,
		tail: nil,
	}
}
