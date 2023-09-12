// Мемоизация в рекурсии . Оптимизация расчета чисел Фибоначчи
package main

import "fmt"

func memo_fib(n int) int {
	res := map[int]int{}
	if n < 2 {
		return n
	}
	if value, ok := res[n]; ok {
		return value
	}
	res[n] = memo_fib(n-1) + memo_fib(n-2)
	return res[n]
}

func main() {
	num := 0
	fmt.Println("Введите число для расчета числа Фибоначчи:")
	fmt.Scanln(&num)
	fmt.Println("Число Фибоначчи равно: ", memo_fib(num))
}
