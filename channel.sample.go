package main

import (
	"fmt"
	"time"
)

// func worker(doen chan bool) {
// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")
// 	done <- true
// }
func main() {
	// done := make(chan bool, 1)
	// go worker()
	// <-done
	c1 := make(chan string)
	c2 := make(chan string)
	// c3 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		// fmt.Println("go 2")
		c1 <- "two"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("go 1")
		c2 <- "one"
	}()
	select {
	case msg1 := <-c1:
		if msg1 == "two" {
			fmt.Println("received", msg1)
		}
		// default:
		// 	fmt.Println("no 2")
	}
	select {
	case msg1 := <-c1:
		fmt.Println("received2", msg1)
	case msg2 := <-c2:
		fmt.Println("received2", msg2)
	default:
		fmt.Println("no 2")

	}
}
