//Exercício de concorrência - condição de corrida
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const quantidadeDeGorotines = 100
var c = 0

func main() {
	fmt.Println(c)
	cria(quantidadeDeGorotines)
	wg.Wait()
	fmt.Println(c)
}

func cria(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func () {
			s := c
			runtime.Gosched()
			s++
			c = s
			wg.Done()
		}()
	}
}