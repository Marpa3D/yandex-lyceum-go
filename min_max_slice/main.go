// ДЗ. min, max в слайсе
package main

import "fmt"

func FindMinMaxInSlice(slice []int) (int, int) {
	if len(slice) == 0 || slice == nil {
		return 0, 0
	}
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
	s1 := []int{} // пример пустого слайса.

	fmt.Println(FindMinMaxInSlice(s))
	fmt.Println(FindMinMaxInSlice(s1)) // Ошибка - error: index out of range [0] with length 0

}
