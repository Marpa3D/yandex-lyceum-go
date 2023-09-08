// ДЗ. Пять шагов. Реверс слайса(среза)
package main

import "fmt"

func ReverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ReverseSlice(s))
}
