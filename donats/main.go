// Яндекс. Лицей. ДЗ. Пончики
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		} else {
			sum += i
		}
	}
	fmt.Println(sum)
}
