package main

import "fmt"

const nome string = "Henrique"

func main() {

	fmt.Println(nome)

	const n = 50000
	n = 1

	fmt.Println(n)
}