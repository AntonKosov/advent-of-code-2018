package list

import (
	slist "container/list"
)

type Circular[T any] struct {
	items slist.List
}

func (c *Circular[T]) Add(value T) *Element[T] {
	return newElement(c, c.items.PushBack(value))
}

func (c *Circular[T]) Remove(element *Element[T]) T {
	return c.items.Remove(element.element).(T)
}

func (c *Circular[T]) InsertAfter(value T, element *Element[T]) *Element[T] {
	return newElement(c, c.items.InsertAfter(value, element.element))
}

type Element[T any] struct {
	list    *Circular[T]
	element *slist.Element
}

func newElement[T any](list *Circular[T], element *slist.Element) *Element[T] {
	return &Element[T]{
		list:    list,
		element: element,
	}
}

func (e *Element[T]) Value() T {
	return e.element.Value.(T)
}

func (e *Element[T]) Next() *Element[T] {
	next := e.element.Next()
	if next == nil {
		next = e.list.items.Front()
	}

	return newElement(e.list, next)
}

func (e *Element[T]) Prev() *Element[T] {
	prev := e.element.Prev()
	if prev == nil {
		prev = e.list.items.Back()
	}

	return newElement(e.list, prev)
}
