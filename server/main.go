package main

import (
	"fmt"
)

func main() {
	go servehttp()
	
	signal := make(chan int, 0)
	<-signal
	fmt.Println("get shut signal, program will be shut")
}
