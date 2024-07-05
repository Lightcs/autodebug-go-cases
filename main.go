package main

import (
	"fmt"
	"main/calculate"
	"math/rand"
	"sync"
)

func generateRandomArray(n, rangeL, rangeR int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(rangeR-rangeL+1) + rangeL // 生成随机数
	}
	return arr
}

func main() {
	correctTimes := map[bool]int{true: 0, false: 0}
	uniqNumbers := map[int]int{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			arr := generateRandomArray(100, 0, 100000)
			for _, v := range arr {
				uniqNumbers[v]++
			}
			calculate.QSort(arr, 0, len(arr)-1)
			for i := 1; i < len(arr); i++ {
				if arr[i-1] > arr[i] {
					correctTimes[false]++
					return
				}
			}
			correctTimes[true]++
		} ()
	}
	wg.Wait()
	fmt.Println(correctTimes)
	fmt.Printf("got %d unique rand numbers\n", len(uniqNumbers))
}
