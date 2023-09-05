// ДЗ.Равенство
package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)

	switch {
	case num1 < 0 || num2 < 0 || num3 < 0:
		fmt.Println("Некорректный ввод")
	case num1 == num2 && num2 == num3 && num1 == num3:
		fmt.Println("Все числа равны")
	case num1 == num2 || num2 == num3 || num1 == num3:
		fmt.Println("Два числа равны")
	case num1 != num2 && num2 != num3 && num1 != num3:
		fmt.Println("Все числа разные")
	}
}
