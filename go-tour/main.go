package main

import (
	"fmt"
	"os"
)

func main() {
	result,err := Division(6,0)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}