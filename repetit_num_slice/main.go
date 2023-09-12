// Подсчет повторений в слайсе
package main

import "fmt"

// countOcc - функция подсчета повторений в слайсе
func countOcc(nums []int) map[int]int {

	res := make(map[int]int)

	for _, num := range nums {
		res[num]++
	}
	return res
}

func main() {
	n := []int{1, 3, 4, 1, 8, 3, 4, 8, 8}

	result := countOcc(n)

	fmt.Println("Результат подсчета повторений:")

	for num, count := range result {
		fmt.Printf("%d встерачается %d раза\n", num, count)
	}
}
