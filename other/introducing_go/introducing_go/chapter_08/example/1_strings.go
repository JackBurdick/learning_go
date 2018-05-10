package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains(string, substring string) bool
	fmt.Println(strings.Contains("ATGTTCTGTGCCACA", "TTC"))

	// Check prefix
	fmt.Println(strings.HasPrefix("ATGTTCTGTGCCACA", "ATG"))

	// Check suffix
	fmt.Println(strings.HasSuffix("ATGTTCTGTGCCTCA", "TCA"))

	// Find index - 0 based
	fmt.Println(strings.Index("XXXXXATGXXXXXXXXX", "ATG"))

	// Join two strings
	fmt.Println(strings.Join([]string{"AAAA", "BBBBB"}, ""))

	// Join two strings
	fmt.Println(strings.Join([]string{"AAAA", "BBBBB"}, "-"))

	// Repeat a string
	fmt.Println(strings.Repeat("Marsha", 3))

	// Split a string
	fmt.Println(strings.Split("a-b-c-d-e-f", "-"))

	// Convert to lower and upper
	fmt.Println(strings.ToLower("HiHowAreYou"))
	fmt.Println(strings.ToUpper("HiHowAreYou"))

	// dealing with bytes
	arr := []byte("test")
	fmt.Println(arr)

	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)

}
