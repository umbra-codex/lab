package iteration

import "strings"

// Repeat returns a new string consisting of times copies of character.
func Repeat(character string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
