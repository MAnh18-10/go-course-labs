package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Задание 2: знак числа
func sign(n int) string {
	switch {
	case n > 0:
		return "Positive"
	case n < 0:
		return "Negative"
	default:
		return "Zero"
	}
}

// Задание 4: длина строки в символах
func strLen(s string) int {
	return utf8.RuneCountInString(s)
}

// Задание 5: структура Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Задание 6: среднее двух int
func averageInt(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Задание 1: чётное или нечётное
	fmt.Print("Введите число: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Ошибка: это не число")
		return
	}

	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	// Задание 2
	fmt.Println("Знак числа:", sign(n))

	// Задание 3: числа 1..10
	fmt.Print("Числа от 1 до 10: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Задание 4
	fmt.Println("Длина строки 'Привет!':", strLen("Привет!"))

	// Задание 5
	r := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Площадь прямоугольника: %.2f\n", r.Area())

	// Задание 6
	fmt.Printf("Среднее значение 4 и 9: %.2f\n", averageInt(4, 9))
}
