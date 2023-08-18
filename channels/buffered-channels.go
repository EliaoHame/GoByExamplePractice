package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	go func() {
		ch <- 3
		fmt.Println("Successfully sent!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
