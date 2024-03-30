package main

import "fmt"

func media(a, b float64) float64 {
	soma_pesos := 11.0
	media := ((a * 3.5) + (b * 7.5)) / soma_pesos
	return media
}

func main() {

	var a, b float64
	fmt.Println("Digite a primeira nota: ")
	fmt.Scanln(&a)
	fmt.Println("Digite a segunda nota: ")
	fmt.Scanln(&b)

	resultado := media(a, b)
	fmt.Printf("MEDIA = %.5f\n", resultado)
}
