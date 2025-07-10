package main

import "fmt"

func mainMaps() {
	idade := map[string]int{}
	idade["steph"] = 28
	idade["ryan"] = 18
	fmt.Println(idade)
	fmt.Println(idade["steph"])
	fmt.Println(idade["ryan"])


	anoNasc := map[string]int{
		"steph": 1995,
		"ryan": 1983,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["steph"])
	fmt.Println(anoNasc["ryan"])
	anoNasc["joão"] = 2017
	fmt.Println(anoNasc)
}