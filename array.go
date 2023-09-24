package main

import "fmt"

type ArrayOperations interface {
	serach(int32) (int32, bool)
	sort(*[]int32)
	merge([]int32, []int32) []int32
	copy(destination, source []int32)
	reverse([]int32) []int32
}

func Array() {
	// Declaration
	Array1 := [5]int32{1, 2, 3, 4, 5}
	var Array2 = [3]int32{0, 5, 6}
	var Array3 = [1]int32{9}
	fmt.Println(Array1, Array2, Array3)
}
