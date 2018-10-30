package main

import "fmt"

type types struct {
	myInt    int
	myString string
}

func main() {
	fmt.Println(types{1, "asd"})

	test := types{1123, "adsasd"}
	fmt.Println(test.myString)
}
