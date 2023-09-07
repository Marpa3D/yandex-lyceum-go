// Функции. ДЗ. Сумма ряда
package main

import "fmt"

func CalculateSeriesSum(n int) float64 {
	//написать программу,
	//которая по заданному числу n
	//будет вычислять значение следующей последовательности: 1 + 1/2 + 1/3 + ... + 1/n
	var value float64
	for i := 0; i < n; i++ {
		value += 1 / (float64(i) + 1)
	}
	return value
}

func main() {
	fmt.Println(CalculateSeriesSum(6))

}
