package main

import "fmt"

// fibonacci generates fibonacci numbers and sends them to channel c
// It stops when it receives a signal on the quit channel
func fibonacci(c, quit chan int) {
	x, y := 0, 1  // Starting values for fibonacci sequence
	
	// Infinite loop - will be terminated by the quit signal
	for {
		select {
		// TRY TO SEND the current fibonacci number to channel c
		// This case will ONLY succeed if someone is ready to RECEIVE from channel c
		// If no one is receiving (channel is blocked), this case won't execute
		case c <- x:
			// We successfully sent x! Now update for next fibonacci number
			x, y = y, x+y
			
		// TRY TO RECEIVE from the quit channel
		// This case will ONLY succeed if someone has SENT to the quit channel
		// If no one has sent (channel is empty), this case won't execute
		case <-quit:
			// We received the quit signal! Time to stop
			fmt.Println("quit")
			return  // Exit the function completely
		}
		// The select statement will keep trying these two operations
		// It blocks until ONE of them can succeed
	}
}

func main() {
	// Create two UNBUFFERED channels
	// Unbuffered channels require perfect synchronization:
	// - A send operation BLOCKS until someone is ready to receive
	// - A receive operation BLOCKS until someone is ready to send
	c := make(chan int)    // Channel for sending fibonacci numbers
	quit := make(chan int) // Channel for sending quit signal
	
	// Start a new goroutine (concurrent execution)
	go func() {
		// This goroutine will RECEIVE 10 fibonacci numbers
		for i := 0; i < 10; i++ {
			// TRY TO RECEIVE from channel c
			// This will BLOCK until fibonacci() tries to send (c <- x)
			fmt.Println(<-c)
		}
		// After receiving 10 numbers, send quit signal
		// TRY TO SEND 0 to quit channel
		// This tells fibonacci() to stop generating numbers
		quit <- 0
		// Goroutine exits after this
	}()
	
	// Call fibonacci in the MAIN goroutine (not as a goroutine)
	// This will block here until fibonacci() returns (after receiving quit signal)
	fibonacci(c, quit)
	
	// When fibonacci() returns, main() ends and program exits
}