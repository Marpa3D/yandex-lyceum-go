// Яндекс. Лицей. ДЗ. За бабушку
package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
