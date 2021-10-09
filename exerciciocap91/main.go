//concorrencia
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Exercício de concorrência.")
	go say()
	go tell()
	wg.Wait()
}

func say() {
	fmt.Println("Hello, from say")
	wg.Done()
}

func tell() {
	fmt.Println("Hello, from tell")
	wg.Done()
}