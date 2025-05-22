package main

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func main() {
	resultado := Somar(3, 5)
	fmt.Println("Resultado:", resultado)
}
