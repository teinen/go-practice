package main

import (
	"fmt"
)

func main() {
	hello := "こんにちは世界"
	runeValue := []rune(hello)

	// get size
	size := len(runeValue)

	// count string size
	fmt.Printf("%s is %d characters\n", hello, size)

	// get each unicode
	for i := 0; i < size; i++ {
		fmt.Printf("%#U\n", runeValue[i])
	}

	fmt.Printf("\n")

	// get each unicode with position
	for pos, value := range runeValue {
		fmt.Printf("%#U starts at byte position %d\n", value, pos)
	}
}
