// package stringutils — пакет для работы со строками.
package stringutils

// Reverse переворачивает строку с учётом Unicode.
// "привет" → "тевирп"
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
