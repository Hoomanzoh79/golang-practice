package main

import (
	"fmt"
	"sync"
)

var counter int 
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	// This is my channel reciever Goroutine
	// The <- points out of the channel meaning we want to recieve
	// This would work the same without the function arguement but wouldn't be as clear
	go func (ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	// This is my channel sender Goroutine
	// The <- points towards the channel meaning we want to send
	// This would work the same without the function arguement but wouldn't be as clear
	go func (ch chan<- int){
		ch <- 43
		ch <- 26
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}