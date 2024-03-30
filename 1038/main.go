package main

import "fmt"

func total(a, b int) float64 {

	switch {
	case a == 1:
		return (float64(b) * 4.00)
	case a == 2:
		return (float64(b) * 4.50)
	case a == 3:
		return (float64(b) * 5.00)
	case a == 4:
		return (float64(b) * 2.00)
	case a == 5:
		return (float64(b) * 1.50)
	}
	return 0
}

func main() {

	var a, b int

	fmt.Println("Digite o código do pedido:")
	fmt.Scan(&a)
	fmt.Println("Digite a quantidade do pedido:")
	fmt.Scan(&b)

	result_total := total(a, b)
	fmt.Printf("Total: R$ %.2f\n", result_total)
}
