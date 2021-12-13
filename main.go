package main

import (
	"algos/sort"
	"fmt"
)

func main() {

	//qsort
	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}
	sort.Qsort(arr)
	fmt.Println(arr)
}
