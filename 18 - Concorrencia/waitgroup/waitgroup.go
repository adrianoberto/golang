package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgroup sync.WaitGroup

	waitgroup.Add(4)

	go func() {
		escrever("Programando em go 1")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em go 2")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em go 3")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programando em go 4")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
