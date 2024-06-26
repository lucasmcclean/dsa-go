package qsort

import (
	"fmt"
	"testing"

	"github.com/ljmcclean/dsa-go/utils"
)

func TestSortEmpty(t *testing.T) {
	slice := utils.GenSliceEmpty(5)
	Sort(slice)
	fmt.Println(slice)
	fmt.Println("Completed Empty")
}

func TestSortNil(t *testing.T) {
	slice := utils.GenSliceNil()
	Sort(slice)
	fmt.Println("Completed Nil")
}

func TestSortInt(t *testing.T) {
	slice := utils.GenSliceInt(100, 10)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Integer")
}

func TestSortIntAsc(t *testing.T) {
	slice := utils.GenSliceIntAsc(1000)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Ascending")
}

func TestSortIntDes(t *testing.T) {
	slice := utils.GenSliceIntDes(1000)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Descending")
}

func TestSortString(t *testing.T) {
	slice := utils.GenSliceStr(1000, 5)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed String")
}

func TestSortStringEmpty(t *testing.T) {
	slice := utils.GenSliceStr(1000, 0)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Empty String")
}

func TestSortFloat(t *testing.T) {
	slice := utils.GenSliceFloat(1000, 10)
	Sort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Float")
}

// Paralell Variations
func TestParallelSortEmpty(t *testing.T) {
	slice := utils.GenSliceEmpty(5)
	PSort(slice)
	fmt.Println("Completed Parallel Empty")
}

func TestParallelSortNil(t *testing.T) {
	slice := utils.GenSliceNil()
	PSort(slice)
	fmt.Println("Completed Parallel Nil")
}

func TestParallelSortInt(t *testing.T) {
	slice := utils.GenSliceInt(100, 10)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel Integer")
}

func TestParallelSortIntAsc(t *testing.T) {
	slice := utils.GenSliceIntAsc(1000)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel Ascending")
}

func TestParallelSortIntDes(t *testing.T) {
	slice := utils.GenSliceIntDes(1000)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel Descending")
}

func TestParallelSortString(t *testing.T) {
	slice := utils.GenSliceStr(1000, 5)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel String")
}

func TestParallelSortStringEmpty(t *testing.T) {
	slice := utils.GenSliceStr(1000, 0)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel Empty String")
}

func TestParallelSortFloat(t *testing.T) {
	slice := utils.GenSliceFloat(1000, 10)
	PSort(slice)
	if !utils.IsSorted(slice) {
		t.Fatalf("Failed to sort")
	}
	fmt.Println("Completed Parallel Float")
}
