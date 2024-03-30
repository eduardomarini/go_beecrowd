package main

import "fmt"

func validar_valores(a, b, c, d int) {
	if b > c && d > a && (c+d) > (a+b) && c > 0 && d > 0 && a%2 == 0 {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
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

	validar_valores(a, b, c, d)

}
