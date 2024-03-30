package main

import "fmt"

func valores(a, b, c, d int) int {
	return ((a * b) - (c * d))
}

func main() {

	var a, b, c, d int

	fmt.Println("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Println("Digite o valor de B: ")
	fmt.Scan(&b)
	fmt.Println("Digite o valor de C: ")
	fmt.Scan(&c)
	fmt.Println("Digite o valor de D: ")
	fmt.Scan(&d)

	resultado := valores(a, b, c, d)
	fmt.Println("DIFERENCA =", resultado)

}
