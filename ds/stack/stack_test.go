package stack

import (
	"testing"

	"github.com/ljmcclean/dsa-go/types"
)

func TestStackWithInts(t *testing.T) {
	stack := New(10, 1, 2, 3, 4, 5)

	got, _ := stack.Peek()
	want := 5
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}

	got, _ = stack.Pop()
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}

	stack.Push(42)
	got, _ = stack.Peek()
	want = 42
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestStackWithStrings(t *testing.T) {
	stack := New(10, "One", "Two", "Three", "Four", "Five")

	got, _ := stack.Peek()
	want := "Five"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}

	got, _ = stack.Pop()
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}

	stack.Push("Forty-Two")
	got, _ = stack.Peek()
	want = "Forty-Two"
	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func TestEmptyStack(t *testing.T) {
	stack := New[int](10)

	_, got := stack.Peek()
	want := types.ErrEmptySlice
	if got != want {
		t.Errorf("Got %e, wanted %e", got, want)
	}

	stack.Push(1)
	_, got = stack.Peek()
	want = nil
	if got != want {
		t.Errorf("Got %e, wanted %e", got, want)
	}
}
