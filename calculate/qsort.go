package calculate

func QSort(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr, low, high)
		QSort(arr, low, pivot-1)
		QSort(arr, pivot+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
