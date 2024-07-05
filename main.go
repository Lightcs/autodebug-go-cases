package main

import (
	"fmt"
	"main/calculate"
	"math/rand"
)

func generateRandomArray(n, rangeL, rangeR int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(rangeR-rangeL+1) + rangeL // 生成随机数
	}
	return arr
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(calculate.Fibonacci(i))
	}
}
