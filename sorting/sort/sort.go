package sort

// QuickSort algorithm.
func QuickSort(arr []int, low int, high int, asc bool) ([]int, int) {
	if low >= high {
		return arr, low
	}

	pivotIndex := high // int((high - low)/2) + low
	pivot := arr[pivotIndex]

	small := low - 1
	for i := low; i < high; i++ {
		if asc {
			if arr[i] < pivot {
				small++
				currentV := arr[small]
				arr[small] = arr[i]
				arr[i] = currentV
			}
		} else {
			if arr[i] > pivot {
				small++
				currentV := arr[small]
				arr[small] = arr[i]
				arr[i] = currentV
			}

		}
	}

	small++
	currentV := arr[small]
	arr[small] = arr[pivotIndex]
	arr[pivotIndex] = currentV
	pivotIndex = small

	_, pivotIndex = QuickSort(arr, low, pivotIndex-1, asc)
	_, pivotIndex = QuickSort(arr, pivotIndex+1, high, asc)

	return arr, pivotIndex
}
