package isort

import (
	"fmt"
	"testing"

	"github.com/ljmcclean/dsa-go/utils"
)

func TestSortEmpty(t *testing.T) {
	slice := utils.GenSliceEmpty(5)
	Sort(slice)
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
