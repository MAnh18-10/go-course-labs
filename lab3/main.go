package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lab3/mathutils"
	"lab3/stringutils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Задание 2: факториал введённого числа
	fmt.Print("Введите число для факториала: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Ошибка: это не число")
		return
	}
	fmt.Printf("Факториал %d = %d\n", n, mathutils.Factorial(n))

	// Задание 3: переворот строки
	fmt.Print("Введите строку для переворота: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println("Перевёрнутая строка:", stringutils.Reverse(s))

	// Задание 4: массив из 5 целых чисел
	var arr [5]int
	for i := range arr {
		arr[i] = (i + 1) * 10
	}
	fmt.Println("Массив:", arr)

	// Задание 5: срез — добавление и удаление
	slice := arr[:]
	slice = append(slice, 60, 70)
	fmt.Println("После добавления:", slice)

	slice = append(slice[:2], slice[3:]...)
	fmt.Println("После удаления индекса 2:", slice)

	// Задание 6: самая длинная строка
	strs := []string{"Go", "программирование", "код", "лабораторная"}
	longest := ""
	for _, v := range strs {
		if len([]rune(v)) > len([]rune(longest)) {
			longest = v
		}
	}
	fmt.Println("Самая длинная строка:", longest)
}
