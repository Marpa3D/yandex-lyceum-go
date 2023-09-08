// ДЗ. Покупки. Сумма элементов слайса
package main

import "fmt"

func SumOfSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(SumOfSlice(slice))
}
