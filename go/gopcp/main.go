package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var mutex sync.RWMutex
	var in int = 10
	fmt.Println("lock the lock")
	go func(){
		mutex.Lock()
		in = in+100
		fmt.Println("hello2",in)
		time.Sleep(time.Second*10)
		fmt.Println("hello4",in)
		mutex.Unlock()
		fmt.Println("hello3",in)
	}()
	time.Sleep(time.Second)
	go func(){
		mutex.Lock()
		in = in+100
		fmt.Println("hello1",in)
		mutex.Unlock()
	}()
	time.Sleep(time.Second*20)
}