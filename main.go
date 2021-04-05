package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person1() {
	Printer("hello")
	ch <- 666
}

func Person2() {
	<-ch
	Printer("world")
}

func main() {
	go Person1()
	go Person2()
	for {

	}
}
