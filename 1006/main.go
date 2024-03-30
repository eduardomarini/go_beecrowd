package main

import (
	"fmt"
)

func validarNotas(a, b, c float64) bool {
	if a < 0 || a > 10 {
		return false
	}
	if b < 0 || b > 10 {
		return false
	}
	if c < 0 || c > 10 {
		return false
	}
	return true
}

func media(a, b, c float64) float64 {

	soma_pesos := 2.0 + 3.0 + 5.0
	media := ((a * 2.0) + (b * 3.0) + (c * 5.0)) / soma_pesos

	return media
}

func main() {

	var a, b, c float64

	fmt.Println("Digite a primeira nota: ")
	fmt.Scanln(&a)
	if !validarNotas(a, b, c) {
		fmt.Println("Nota 1 invalida, digite uma nota entre 0 e 10.")
		return
	}
	fmt.Println("Digite a segunda nota: ")
	fmt.Scanln(&b)
	if !validarNotas(a, b, c) {
		fmt.Println("Nota 2 invalida, digite uma nota entre 0 e 10.")
		return
	}
	fmt.Println("Digite a terceira nota: ")
	fmt.Scanln(&c)
	if !validarNotas(a, b, c) {
		fmt.Println("Nota 3 invalida, digite uma nota entre 0 e 10.")
		return
	}

	resultado := media(a, b, c)

	fmt.Printf("MEDIA = %.1f\n", resultado)

}
