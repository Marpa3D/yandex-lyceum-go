// Яндекс.Лицей. ДЗ. Перестановки букв. Факториал
package main

import "fmt"

func main() {
	var n int
	f := 1
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		f = f * i
	}
	fmt.Println(f)
}
