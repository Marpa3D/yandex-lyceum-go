// ДЗ. Много звезд
package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	switch {
	case num < 0:
		fmt.Println("Некорректный ввод")
	case num < 10:
		fmt.Println("Число меньше 10")
	case num < 100:
		fmt.Println("Число меньше 100")
	case num < 1000:
		fmt.Println("Число меньше 1000")
	case num >= 1000:
		fmt.Println("Число больше или равно 1000")
	}
}
