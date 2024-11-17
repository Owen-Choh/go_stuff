package sort

import (
	// "fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		asc bool
	}{
		{[]int{4, 1, 9, 2}, []int{1, 2, 4, 9}, true},
		{[]int{1, 4, 1, 9, 2}, []int{1, 1, 2, 4, 9}, true},
		{[]int{4, 1, 9, 5, 2}, []int{1, 2, 4, 5, 9}, true},
		{[]int{4, 1}, []int{1, 4}, true},
		{[]int{4, 1, 9, 2}, []int{9,4,2,1}, false},
		{[]int{1, 4, 1, 9, 2}, []int{9,4,2,1,1}, false},
		{[]int{4, 1, 9, 5, 2}, []int{9,5,4,2,1}, false},
		{[]int{4, 1}, []int{4,1}, false},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1)
		go func() {
			defer wg.Done()

			output, _ := QuickSort(test.input, 0, len(test.input)-1, test.asc)
			for i, v := range output {
				if v != test.expected[i] {
					t.Logf("Test failed: expected %v, got %v", test.expected, output)
					t.Fail()
					return
				}
			}
			t.Logf("test input passed %v", test.expected)
		}()
	}

	wg.Wait()

	t.Log("test over")
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		asc bool
	}{
		{[]int{4, 1, 9, 2}, []int{1, 2, 4, 9}, true},
		{[]int{1, 4, 1, 9, 2}, []int{1, 1, 2, 4, 9}, true},
		{[]int{4, 1, 9, 5, 2}, []int{1, 2, 4, 5, 9}, true},
		{[]int{4, 1}, []int{1, 4}, true},
		{[]int{4, 1, 9, 2}, []int{9,4,2,1}, false},
		{[]int{1, 4, 1, 9, 2}, []int{9,4,2,1,1}, false},
		{[]int{4, 1, 9, 5, 2}, []int{9,5,4,2,1}, false},
		{[]int{4, 1}, []int{4,1}, false},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1)
		go func() {
			defer wg.Done()

			output := MergeSort(test.input, 0, len(test.input)-1, test.asc)
			for i, v := range output {
				if v != test.expected[i] {
					t.Logf("Test failed: expected %v, got %v", test.expected, output)
					t.Fail()
					return
				}
			}
			t.Logf("test input passed %v", test.expected)
		}()
	}

	wg.Wait()

	t.Log("test over")
}
func TestMergeSortParallel(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		asc bool
	}{
		{[]int{4, 1, 9, 2}, []int{1, 2, 4, 9}, true},
		{[]int{1, 4, 1, 9, 2}, []int{1, 1, 2, 4, 9}, true},
		{[]int{4, 1, 9, 5, 2}, []int{1, 2, 4, 5, 9}, true},
		{[]int{4, 1}, []int{1, 4}, true},
		{[]int{4, 1, 9, 2}, []int{9,4,2,1}, false},
		{[]int{1, 4, 1, 9, 2}, []int{9,4,2,1,1}, false},
		{[]int{4, 1, 9, 5, 2}, []int{9,5,4,2,1}, false},
		{[]int{4, 1}, []int{4,1}, false},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1)
		go func() {
			defer wg.Done()

			output := MergeSortParallel(test.input, 0, len(test.input)-1, test.asc)
			for i, v := range output {
				if v != test.expected[i] {
					t.Logf("Test failed: expected %v, got %v", test.expected, output)
					t.Fail()
					return
				}
			}
			t.Logf("test input passed %v", test.expected)
		}()
	}

	wg.Wait()

	t.Log("test over")
}

const randSeed = 10
func BenchmarkMergeSortParallel(b *testing.B) {
	// results: 20558 iterations, 50822 ns/op time taken, 81361 B/op memory, 2001 allocs/op memory allocs
	rand.NewSource(randSeed)
	input := rand.Perm(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeSortParallel(input,0,len(input)-1,true)
	}
}
func BenchmarkMergeSort(b *testing.B) {
	// results: 17902 iterations, 64154 ns/op time taken, 81152 B/op memory, 1998 allocs/op memory allocs
	rand.NewSource(randSeed)
	input := rand.Perm(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MergeSort(input,0,len(input)-1,true)
	}
}
