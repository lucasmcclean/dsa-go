package qsort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	amount := 100000000
	slice1 := make([]int, 0, amount)
	slice2 := make([]int, 0, amount)
	for i := 0; i < amount; i++ {
		random := rand.Int()
		slice1 = append(slice1, random)
		slice2 = append(slice2, random)
	}

	start := time.Now()
	Sort(slice1)
	duration := time.Since(start)
	fmt.Println("Sort:  ", duration)

	start = time.Now()
	PSort(slice2)
	duration = time.Since(start)
	fmt.Println("PSort: ", duration)

	for i := 1; i < amount; i++ {
		if slice1[i] < slice1[i-1] {
			t.Fatal("Failed to sort 1")
		}
		if slice2[i] < slice2[i-1] {
			t.Fatal("Failed to sort 2")
		}
	}
}
