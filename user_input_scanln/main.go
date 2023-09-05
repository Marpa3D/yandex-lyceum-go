package main

import "fmt"

func main() {
	var age int
	fmt.Println("Введите ваш возраст:")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Printf("Ошибка %s", err)
	}
	fmt.Printf("Ваш возраст: %d лет.", age)
	fmt.Printf("\nАдрес, где лежит считанное значение age: %p", &age)
}
