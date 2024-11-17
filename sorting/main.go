package main

import (
	"fmt"
	"github.com/Owen-Choh/go_stuff/sorting/sort"
)

func main()  {
	data := []int{1,7,4,2,6,5}
	result := sort.MergeSort(data,0,len(data)-1, true)
	fmt.Println("sorted data:",result)
}