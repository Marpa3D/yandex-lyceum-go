// Яндекс.Лицей. ДЗ. Нечетная сторона
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	if n < 0 {
		fmt.Println("Некорректный ввод")
	} else {
		fmt.Println(sum)
	}
}
