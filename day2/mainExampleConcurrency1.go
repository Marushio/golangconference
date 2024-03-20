package main

import "fmt"
//T1
func main() {
	canal := make(chan string)

	//T2
	go func(){
		canal <- "Golang Conference! - Vindo da T2"
	}()

	//T1
	msg := <-canal // se o canal siver alguma coisa, joga em MSG
	fmt.Println(msg)

	ch := make(chan int) 
	go publish(ch)
	consume(ch)
}

func publish(ch chan int){
	for i := 0; i < 10; i++ {
		ch <- i //nunca enche a var com mais de uma variavel espera o consume consumir o ch
	}
	close(ch)
}

func consume(ch chan int){
	for x := range ch { //esvazia canal.
		fmt.Println("Mensagem procesada", x) //
	}
}