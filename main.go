package main

import "fmt"

// Test constants
const (
	FibLength = 10
)

func main() {
	runFibo(FibLength)
}

func runFibo(l int) {
	num1 := 0
	num2 := 1
	numSum := num1 + num2
	fmt.Println(numSum)
	for i := 0; i < l; i++ {
		numSum = num1 + num2
		num1 = num2
		num2 = numSum
		fmt.Println(numSum)
	}
}
