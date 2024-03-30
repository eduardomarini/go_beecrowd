package main

import (
	"fmt"
)

func notas(valor int) {

	notas := [7]int{100, 50, 20, 10, 5, 2, 1}

	for _, notas := range notas {

		quantidade := valor / notas
		valor %= notas

		fmt.Printf("%d nota(s) de R$ %d,00\n", quantidade, notas)
	}
}

func main() {

	var valor int

	fmt.Println("Digite o valor: ")
	fmt.Scan(&valor)

	notas(valor)

}
