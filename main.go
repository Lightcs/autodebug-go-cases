package main

import (
	"fmt"
	"main/calculate"
)

func main() {
	arr := []int{5, 2, 7, 1, 9}
	fmt.Println("原始数组:", arr)

	calculate.QSort(arr, 0, len(arr)-1)
	fmt.Println("排序后的数组:", arr)

	expected := []int{1, 2, 5, 7, 9}
	match := true
	for i, v := range arr {
		if v != expected[i] {
			match = false
			break
		}
	}
	if !match {
		panic("排序失败")
	}
}
