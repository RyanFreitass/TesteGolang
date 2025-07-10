package main

import "fmt"

// STRUCTS
//   Forma de criar sua própria estrutura de dados
//	 Personalizar de acordo com a sua necessidade
//	 Podemos usar vários tipos diferentes

// type <nome da estrutura> struct { <CAMPOS> }
type Pessoa struct {
	Nome string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Ryan", 18})
	fmt.Println(Pessoa{Nome: "João", Idade: 58})
	fmt.Println(Pessoa{Nome: "Steph", Idade: 20})

	p1 := Pessoa{Nome: "Bob", Idade: 12}
	//fmt.Println(p1.Nome)
	//fmt.Println(p1.Idade)

	//p1.Idade = 3
	//fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 5}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// structs + map
	//alunos := map[string][]Pessoa{}
	//alunos["programação"] = pessoas
	//fmt.Println(alunos)

	

	//var alunos = map[string][]Pessoa {
	//	"Programação": {{Nome: "Ryan", Idade: 18}, {Nome: "Steph", Idade: 21}},
	//}
	//fmt.Println(alunos)


	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
