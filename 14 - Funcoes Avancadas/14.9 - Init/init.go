package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando o init")
	n = 10
}

func main() {
	fmt.Println("Executando o main")
	fmt.Println(n)
}
