package sort

import "sync"

// standard QuickSort
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
				arr[small], arr[i] = arr[i], arr[small]
			}
		} else {
			if arr[i] > pivot {
				small++
				arr[small], arr[i] = arr[i], arr[small]
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

// standard mergesort
func MergeSort(arr []int, low int, high int, asc bool) []int {
	if low >= high {
		return arr
	}

	mid := int((high-low)/2) + low

	var left []int = make([]int, mid-low+1)
	copy(left, MergeSort(arr, low, mid, asc)[low:mid+1])

	var right []int = make([]int, high-mid)
	copy(right, MergeSort(arr, mid+1, high, asc)[mid+1:high+1])

	current := low
	j := 0
	i := 0
	for i < len(left) && j < len(right) {
		if asc {
			if left[i] < right[j] {
				arr[current] = left[i]
				i++
			} else {
				arr[current] = right[j]
				j++
			}
		} else {
			if left[i] > right[j] {
				arr[current] = left[i]
				i++
			} else {
				arr[current] = right[j]
				j++
			}
		}
		current++
	}

	for ; i < len(left); i++ {
		arr[current] = left[i]
		current++
	}

	for ; j < len(right); j++ {
		arr[current] = right[j]
		current++
	}

	return arr
}

// mergesort but with some concurrency
func MergeSortParallel(arr []int, low int, high int, asc bool) []int {
	if low >= high {
		return arr
	}

	mid := int((high-low)/2) + low

	var wg sync.WaitGroup

	var left []int = make([]int, mid-low+1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		copy(left, MergeSort(arr, low, mid, asc)[low:mid+1])
	}()

	var right []int = make([]int, high-mid)
	wg.Add(1)
	go func() {
		defer wg.Done()
		copy(right, MergeSort(arr, mid+1, high, asc)[mid+1:high+1])
	}()

	wg.Wait()

	current := low
	j := 0
	i := 0
	for i < len(left) && j < len(right) {
		if asc {
			if left[i] < right[j] {
				arr[current] = left[i]
				i++
			} else {
				arr[current] = right[j]
				j++
			}
		} else {
			if left[i] > right[j] {
				arr[current] = left[i]
				i++
			} else {
				arr[current] = right[j]
				j++
			}
		}
		current++
	}

	for ; i < len(left); i++ {
		arr[current] = left[i]
		current++
	}

	for ; j < len(right); j++ {
		arr[current] = right[j]
		current++
	}

	return arr
}
