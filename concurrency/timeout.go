package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan string,1)
	go func () {
		time.Sleep(2 * time.Second)
		ch1 <- "response 1"
	}()
	select {
	case res := <-ch1:
		fmt.Println(res)
	// this case is ran for this goroutine
	case <- time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}
	ch2 := make(chan string,1)
	go func () {
		time.Sleep(2 * time.Second)
		ch2 <- "response 2"
	}()
	select {
	// this case is ran for this goroutine
	case res := <-ch2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}