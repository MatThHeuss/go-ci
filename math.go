package main

import "fmt"

func main() {

	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Subt(a int, b int) int {
	return a - b
}

func Divide(a int, b int) int {
	return a / b
}
