package hsort

import (
	"math/rand"
	"testing"
)

func TestHeapSort(t *testing.T) {
	amount := 10000
	var slice []int
	for i := 0; i < amount; i++ {
		random := rand.Int()
		slice = append(slice, random)
	}
	Sort(slice)
	for i := 1; i < amount; i++ {
		if slice[i] < slice[i-1] {
			t.Fatal("Failed to sort")
		}
	}
}
