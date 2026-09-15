package main

import (
	"fmt"
	"time"
)

// Задание 5: сумма и разность двух чисел с плавающей запятой
func sumAndDiff(a, b float64) (float64, float64) {
	return a + b, a - b
}

// Задание 6: среднее значение трёх чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	// Задание 1: текущее время и дата
	now := time.Now()
	fmt.Println("Текущая дата и время:", now.Format("2006-01-02 15:04:05"))

	// Задание 2: переменные различных типов
	var i int = 18
	var f float64 = 3.14
	var s string = "Hello, Go!"
	var b bool = true
	fmt.Printf("int: %d, float64: %.2f, string: %s, bool: %t\n", i, f, s, b)

	// Задание 3: краткая форма объявления переменных
	x := 10
	y := 2.5
	name := "Anna"
	flag := false
	fmt.Printf("x=%d, y=%.1f, name=%s, flag=%t\n", x, y, name, flag)

	// Задание 4: арифметические операции
	a, b2 := 18, 2
	fmt.Printf("%d + %d = %d\n", a, b2, a+b2)
	fmt.Printf("%d - %d = %d\n", a, b2, a-b2)
	fmt.Printf("%d * %d = %d\n", a, b2, a*b2)
	fmt.Printf("%d / %d = %d\n", a, b2, a/b2)
	fmt.Printf("%d %% %d = %d\n", a, b2, a%b2)

	// Задание 5
	sum, diff := sumAndDiff(8.5, 1.5)
	fmt.Printf("Сумма: %.2f, Разность: %.2f\n", sum, diff)

	// Задание 6
	fmt.Printf("Среднее значение: %.2f\n", average(3, 9, 18))
}
