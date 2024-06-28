package stack

import (
	"sync"

	"github.com/ljmcclean/dsa-go/types"
)

// Not protected from paralell read/writes. Use the normal stack
// if you are accessing this data concurrently.
type UnsafeStack[T any] struct {
	length   int
	Elements []T
}

func Unsafe[T any](capacity int, elems ...T) (stack *UnsafeStack[T]) {
	stack = &UnsafeStack[T]{
		Elements: make([]T, 0, capacity),
		length:   len(elems),
	}
	stack.Elements = append(stack.Elements, elems...)
	return stack
}

func (stack *UnsafeStack[T]) Push(elem T) {
	stack.Elements = append(stack.Elements, elem)
	stack.length++
}

func (stack *UnsafeStack[T]) Pop() (elem T, err error) {
	if stack.length == 0 {
		return elem, types.ErrEmptySlice
	}
	elem = stack.Elements[len(stack.Elements)-1]
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
	stack.length--
	return elem, nil
}

func (stack *UnsafeStack[T]) Peek() (elem T, err error) {
	if stack.length == 0 {
		return elem, types.ErrEmptySlice
	}
	return stack.Elements[len(stack.Elements)-1], nil
}

func (stack *UnsafeStack[T]) Len() (length int) {
	return stack.length
}

// Implements a RWMutex for safe concurrent access.
type Stack[T any] struct {
	mu    sync.RWMutex
	Stack UnsafeStack[T]
}

func New[T any](capacity int, elems ...T) (stack *Stack[T]) {
	return &Stack[T]{
		Stack: *Unsafe(capacity, elems...),
	}
}

func (s *Stack[T]) Push(elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Stack.Push(elem)
}

func (s *Stack[T]) Pop() (elem T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Stack.Pop()
}

func (s *Stack[T]) Peek() (elem T, err error) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.Stack.Peek()
}

func (s *Stack[T]) Len() (length int) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.Stack.Len()
}
