package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string, s string) {
	for i := 0; ; i++ {
		fmt.Println("seteando")
		if s != "" {
			c <- s
		} else {
			c <- "ping"
		}
	}
}

func printer(c <-chan string) {

	for {
		fmt.Println("y...")
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		case c <- x:
			x, y = y, x+y
		}
	}
}

func main() {
	c := make(chan int, 10)
	quit := make(chan int)
	go func() {
		quit <- 0
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()
	fibonacci(c, quit)
}
