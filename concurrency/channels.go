package main

func sum(s[] int,c chan <- int){
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum
}

// func main() {
// 	s := []int {1,2,3}
// 	ch := make(chan int)
// 	go sum(s[:1],ch)
// 	go sum(s,ch)
// 	x,y := <-ch,<-ch
// 	fmt.Println(x,y)
// }