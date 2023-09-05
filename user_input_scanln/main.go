package main

import "fmt"

func main() {
	var age int
	fmt.Println("Введите Ваш возраст:")
	fmt.Scanln(&age)
	fmt.Printf("Ваш возраст: %d лет.", age)
}
