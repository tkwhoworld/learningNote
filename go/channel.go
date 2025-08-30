package main

import (
	"fmt"
	"time"
)

func dd2() {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("hello,%d\n", <-intChan)
	}
}

func main() {
	dd2()
	strChan := make(chan int, 10)
	go func() {
		for ele := range strChan {
			fmt.Printf("revice message: %d\n", ele)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("send message2: %d\n", i)
			strChan <- i
			time.Sleep(1 * time.Millisecond)
		}
	}()
	close(strChan)

}
