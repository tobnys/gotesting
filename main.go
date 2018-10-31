package main

import "fmt"

// Test constants
const (
	TEST  = 123
	TEST2 = 1234
)

func main() {
	fmt.Println(TEST)
}

func testFunc(a, b int) int {
	return a + b
}
