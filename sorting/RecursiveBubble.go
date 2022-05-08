package sorting

func RecursiveBubble(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
		i++
	}

	RecursiveBubble(arr, size-1)
	return arr
}