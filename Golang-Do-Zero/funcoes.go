package main

import "fmt"

func main() {
	resultado_soma := soma(42, 13)
	fmt.Println(resultado_soma)

	resultado_sub := sub(10, 5)
	fmt.Println(resultado_sub)

	resultados_div := div(10, 2)
	fmt.Println(resultados_div)

	resultados_mult := mult(10, 2)
	fmt.Println(resultados_mult)

	printName := nome("RYAN")
	fmt.Println(printName)

	nome, sobrenome := nomeCompleto("Ryan", "Freitas")
	fmt.Println(nome)
	fmt.Println(sobrenome)
	fmt.Println(nome, sobrenome)

}

// FUNÇÃO COMEÇANDO COM LETRA MINUSCULA: FUNÇÃO PRIVADA
// FUNÇÃO COMEÇANDO COM LETRA MAIUSCULA: FUNÇÃO PÚBLICA

func nomeCompleto(nome string, sobrenome string) (string, string) {
	return nome, sobrenome
}

func nome(nome string) string {
	return nome
}

func soma(x int, y int) int {
	somaDosValores := x + y
	return somaDosValores
}

func sub(x int, y int) int {
	subDosValores := x - y
	return subDosValores
}

func div(x int, y int) int {
	divDosValores := x / y
	return divDosValores
}

func mult(x int, y int) int {
	multDosValores := x * y
	return multDosValores
}