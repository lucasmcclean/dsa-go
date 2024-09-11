package stack

import "sync"

// An UnsafeStack is not protected from concurrent read/writes.
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

func (stack *UnsafeStack[T]) Push(elem ...T) {
	stack.Elements = append(stack.Elements, elem...)
	stack.length += len(elem)
}

func (stack *UnsafeStack[T]) Pop() (elem T) {
	if stack.length == 0 {
		return elem
	}
	elem = stack.Elements[len(stack.Elements)-1]
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
	stack.length--
	return elem
}

func (stack *UnsafeStack[T]) Peek() (elem T) {
	if stack.length == 0 {
		return elem
	}
	return stack.Elements[len(stack.Elements)-1]
}

func (stack *UnsafeStack[T]) Len() (length int) {
	return stack.length
}

// Stack is safe for concurrent read/writes.
type Stack[T any] struct {
	mu    sync.RWMutex
	Stack UnsafeStack[T]
}

func New[T any](capacity int, elems ...T) (stack *Stack[T]) {
	return &Stack[T]{
		Stack: *Unsafe(capacity, elems...),
	}
}

func (s *Stack[T]) Push(elem ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Stack.Push(elem...)
}

func (s *Stack[T]) Pop() (elem T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Stack.Pop()
}

func (s *Stack[T]) Peek() (elem T) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.Stack.Peek()
}

func (s *Stack[T]) Len() (length int) {
	s.mu.RLock()
	s.mu.RUnlock()
	return s.Stack.Len()
}
