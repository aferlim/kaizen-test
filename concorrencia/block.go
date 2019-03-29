package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(3 * time.Second)

	c <- 1 // operacao bloqueante
	c <- 2
	fmt.Println("Passou por aqui")
	c <- 3
	c <- 5
	c <- 6

	fmt.Println("Só depois que for lido")
}

// func main() {
// 	c := make(chan int)

// 	go routine(c)

// 	<-c
// 	<-c
// 	fmt.Println(<-c)

// 	fmt.Println("foi lido")

// 	//fmt.Println(<-c)

// }
