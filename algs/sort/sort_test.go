package sort

import (
	"testing"

	"github.com/ljmcclean/dsa-go/algs"
)

func TestSortInts(t *testing.T) {
	sliceT := algs.GenSliceInt(100, algs.MaxInt, seed)
	var slice []int
	slice = append([]int{}, sliceT...)
	InsertionSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Insertion sort failed to sort integers")
	}
	slice = append([]int{}, sliceT...)
	ShellSort(slice, 10)
	if !algs.IsSorted(slice) {
		t.Error("Shell sort failed to sort integers")
	}
	slice = append([]int{}, sliceT...)
	QuickSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Quick sort failed to sort integers")
	}
	slice = append([]int{}, sliceT...)
	HeapSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Heap sort failed to sort integers")
	}
}

func TestSortStrs(t *testing.T) {
	sliceT := algs.GenSliceStr(100, 5, seed)
	var slice []string
	slice = append([]string{}, sliceT...)
	InsertionSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Insertion sort failed to sort strings")
	}
	slice = append([]string{}, sliceT...)
	ShellSort(slice, 10)
	if !algs.IsSorted(slice) {
		t.Error("Shell sort failed to sort strings")
	}
	slice = append([]string{}, sliceT...)
	QuickSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Quick sort failed to sort strings")
	}
	slice = append([]string{}, sliceT...)
	HeapSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Heap sort failed to sort strings")
	}
}

func TestSortIntsAsc(t *testing.T) {
	sliceT := algs.GenSliceIntAsc(100)
	var slice []int
	slice = append([]int{}, sliceT...)
	InsertionSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Insertion sort failed to sort integers ascending")
	}
	slice = append([]int{}, sliceT...)
	ShellSort(slice, 10)
	if !algs.IsSorted(slice) {
		t.Error("Shell sort failed to sort integers ascending")
	}
	slice = append([]int{}, sliceT...)
	QuickSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Quick sort failed to sort integers ascending")
	}
	slice = append([]int{}, sliceT...)
	HeapSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Heap sort failed to sort integers ascending")
	}
}

func TestSortIntsDes(t *testing.T) {
	sliceT := algs.GenSliceIntDes(100)
	var slice []int
	slice = append([]int{}, sliceT...)
	InsertionSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Insertion sort failed to sort integers descending")
	}
	slice = append([]int{}, sliceT...)
	ShellSort(slice, 10)
	if !algs.IsSorted(slice) {
		t.Error("Shell sort failed to sort integers descending")
	}
	slice = append([]int{}, sliceT...)
	QuickSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Quick sort failed to sort integers descending")
	}
	slice = append([]int{}, sliceT...)
	HeapSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Heap sort failed to sort integers descending")
	}
}

func TestSortFloats(t *testing.T) {
	sliceT := algs.GenSliceFloat(100, 999999.0, seed)
	var slice []float64
	slice = append([]float64{}, sliceT...)
	InsertionSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Insertion sort failed to sort floats(float64)")
	}
	slice = append([]float64{}, sliceT...)
	ShellSort(slice, 10)
	if !algs.IsSorted(slice) {
		t.Error("Shell sort failed to sort floats(float64)")
	}
	slice = append([]float64{}, sliceT...)
	QuickSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Quick sort failed to sort floats(float64)")
	}
	slice = append([]float64{}, sliceT...)
	HeapSort(slice)
	if !algs.IsSorted(slice) {
		t.Error("Heap sort failed to sort floats(float64)")
	}
}

func TestSortEmpty(t *testing.T) {
	sliceT := algs.GenSliceEmpty(100)
	var slice []int
	slice = append([]int{}, sliceT...)
	InsertionSort(slice)
	ShellSort(slice, 10)
	QuickSort(slice)
	HeapSort(slice)
}

func TestSortNil(t *testing.T) {
	sliceT := algs.GenSliceNil()
	var slice []int
	slice = append([]int{}, sliceT...)
	InsertionSort(slice)
	ShellSort(slice, 10)
	QuickSort(slice)
	HeapSort(slice)
}
