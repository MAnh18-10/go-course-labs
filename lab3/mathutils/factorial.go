// package mathutils — пакет с математическими функциями.
package mathutils

// Factorial вычисляет факториал числа n (n >= 0).
// Имя с большой буквы — функция экспортируется в другие пакеты.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
