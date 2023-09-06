// ДЗ. Квадратные корни
package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c, discriminant, x1, x2 float64
	fmt.Scanln(&a, &b, &c)

	discriminant = b*b - 4*a*c

	if discriminant > 0 {
		x1 = (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 = (-b - math.Sqrt(discriminant)) / (2 * a)

		fmt.Printf("%f %f", x2, x1)
	}
	if discriminant == 0 {
		x1 = (-b + math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("%f", x1)
	}
	if discriminant < 0 {
		fmt.Println("0 0")
	}
}

func main() {
	SqRoots()
}
