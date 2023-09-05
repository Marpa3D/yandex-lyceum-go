// Разделительный ноль
package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		fmt.Println("Первое число больше второго")
	} else if num2 > num1 {
		fmt.Println("Второе число больше первого")
	} else if num1 == num2 {
		fmt.Println("Числа равны")
	}
}
