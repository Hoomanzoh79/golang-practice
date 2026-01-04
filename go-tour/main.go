package main

import "fmt"

func main() {
	// pos, neg := Adder(), Adder()
	f := Fibo()
	for i:= 0 ;i<10;i++{
		// fmt.Println(
		// 	pos(i),
		// 	neg(-2*i),
		// )
		fmt.Println(f())
	}
}
//results are 
// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90
