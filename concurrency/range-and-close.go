package main

func fibo(n int,c chan int) {
	x,y := 0,1
	for i := 0;i < n;i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)
}

// func main() {
// 	ch := make(chan int,10)
// 	go fibo(cap(ch),ch)
// 	for i := range ch {
// 		v, ok := <-ch
// 		if ok {
// 			fmt.Println(i,v)
// 		} else {fmt.Println("Channel is closed")}
// 	}
// }
