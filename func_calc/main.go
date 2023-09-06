// Простой калькулятор в функциях
package main

import "fmt"

//func sum(num1, num2 int) int {
//	return num1 + num2
//}
//
//func raz(num1, num2 int) int {
//	return num1 - num2
//}
//
//func mult(num1, num2 int) int {
//	return num1 * num2
//}
//
//func divide(num1, num2 int) (int, int) {
//	integer := num1 / num2
//	remainder := num1 % num2
//
//	return integer, remainder
//}

func calc(n1, n2 float64) (float64, float64, float64, float64) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

func main() {
	var n1, n2 float64
	//fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&n1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&n2)

	sum, raz, multip, divided := calc(n1, n2)
	//resInt, resRemain := divide(n1, n2)

	fmt.Printf("Результат сложения: %0.1f\n", sum)
	fmt.Printf("Результат вычитания: %0.1f\n", raz)
	fmt.Printf("Результат умножения: %0.1f\n", multip)
	fmt.Printf("Результат деления %0.1f\n", divided)
}
