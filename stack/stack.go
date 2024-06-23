package stack

import (
	"sync"

	"github.com/ljmcclean/dsa-go/utils"
)

type UnsafeStack[T any] struct {
	length   int
	elements []T
}

func Unsafe[T any](capacity int, elems ...T) (stack *UnsafeStack[T]) {
	stack = &UnsafeStack[T]{
		elements: make([]T, 0, capacity),
		length:   len(elems),
	}
	stack.elements = append(stack.elements, elems...)
	return stack
}

func (stack *UnsafeStack[T]) Push(elem T) {
	stack.elements = append(stack.elements, elem)
	stack.length++
}

func (stack *UnsafeStack[T]) Pop() (elem T, err error) {
	if stack.length == 0 {
		return elem, utils.ErrEmptySlice
	}
	elem = stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	stack.length--
	return elem, nil
}

func (stack *UnsafeStack[T]) Peek() (elem T, err error) {
	if stack.length == 0 {
		return elem, utils.ErrEmptySlice
	}
	return stack.elements[len(stack.elements)-1], nil
}

func (stack *UnsafeStack[T]) Len() (length int) {
	return stack.length
}

type Stack[T any] struct {
	mu    sync.RWMutex
	stack UnsafeStack[T]
}

func New[T any](capacity int, elems ...T) (stack *Stack[T]) {
	return &Stack[T]{
		stack: *Unsafe(capacity, elems...),
	}
}

func (s *Stack[T]) Push(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stack.Push(elem)
}

func (s *Stack[T]) Pop() (elem T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stack.Pop()
}

func (s *Stack[T]) Peek() (elem T, err error) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.stack.Peek()
}

func (s *Stack[T]) Len() (length int) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.stack.Len()
}
