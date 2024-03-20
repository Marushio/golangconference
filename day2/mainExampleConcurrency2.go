package main

import (
	"fmt"
	"time"
)
func worker(workerId int, data chan int) {
	for x := range data{
		time.Sleep(time.Second)
		fmt.Printf("Worker %d recieve %d\n", workerId, x)
	}
}
func main() {
	ch := make(chan int)
	
	//jeito 1 de fazer os workers rodarem
	//go worker(1, ch)
	//go worker(2, ch)
	//go worker(3, ch)

	//jeito 2 de fazer os workers rodarem
	qtdWorkers := 10	
	for i := 0; i < qtdWorkers; i++{
		go worker(i, ch)
	} 

	for i := 0; i < 100; i++ {
		ch <- i
	}
}
