package study

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{82, 4, 6, 2, 4, 8, 9, 2, 35, 7}
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestQuicklySort(t *testing.T) {
	//arr := []int{82, 4, 6, 2, 4, 8, 9, 2, 35, 7}
	//QuicklySort(arr, 0, len(arr)-1)
	//fmt.Println(arr)
	fmt.Println(CalcPossibilities("ordeals"))
}
