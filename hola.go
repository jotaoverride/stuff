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
			fmt.Println("x")
			x, y = y, x+y
		}
	}
}

// sell := j["status"].(map[string]interface{})["sell"] 
// if sell != nil {
// 	sellImmediatePayment := sell.(map[string]interface{})["immediate_payment"]
// 	if sellImmediatePayment != nil {
// 		sellRequiredField := sellImmediatePayment.(map[string]interface{})["required"]
// 		if sellRequiredField != nil {
// 			sellRequired := sellRequiredField.(bool)
// 		}
// 	} 
// }

func main() {
	c := make(chan int, 10)
	quit := make(chan int)
	go fibonacci(c, quit)
	func() {
		quit <- 0
		fmt.Println("quit")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()
}
