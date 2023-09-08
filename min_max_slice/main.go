// ДЗ. min, max в слайсе
package main

import "fmt"

func RFindMinMaxInSlice(slice []int) (int, int) {
	min, max := slice[0], slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	s := []int{1, 20, 3, 88, 5, 108, 7}

	fmt.Println(RFindMinMaxInSlice(s))
}
