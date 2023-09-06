// Простой калькулятор в функциях
package main

import "fmt"

func sum(num1, num2 int) int {
	return num1 + num2
}

func raz(num1, num2 int) int {
	return num1 - num2
}

func mult(num1, num2 int) int {
	return num1 * num2
}

func divide(num1, num2 int) (int, int) {
	integer := num1 / num2
	remainder := num1 % num2

	return integer, remainder
}

func main() {
	var n1, n2 int

	//fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&n1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&n2)

	resInt, resRemain := divide(n1, n2)

	fmt.Println("Результат сложения: ", sum(n1, n2))
	fmt.Printf("Результат вычитания: %d\n", raz(n1, n2))
	fmt.Printf("Результат умножения: %d\n", mult(n1, n2))
	fmt.Printf("Результат деления %d и остаток %d\n", resInt, resRemain)
}
