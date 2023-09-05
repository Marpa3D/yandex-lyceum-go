// Яндекс.Лицей. ДЗ. Таблица умножения
package main

import "fmt"

func main() {
	var num, multip int
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		// 8 * 1 = 8...
		multip = i * num
		fmt.Printf("%d * %d = %d\n", num, i, multip)
	}
}
