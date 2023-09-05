// ДЗ. Около нуля
package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scanln(&num1, &num2)

	switch {
	case num1 > 0 && num2 > 0:
		fmt.Println("Оба числа положительные")
	case num1 < 0 && num2 < 0:
		fmt.Println("Оба числа отрицательные")
	case num1 < 0 || num2 < 0:
		fmt.Println("Одно число положительное, а другое отрицательное")
	case num1 == 0 || num2 == 0:
		fmt.Println("Одно из чисел равно нулю")
	}
	fmt.Println("")
}
