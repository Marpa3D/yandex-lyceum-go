// Яндекс.Лицей. ДЗ. Кролики Фибоначчи
package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i <= 9; i++ {
		fmt.Println(fibonacci(n + i))
	}
}
