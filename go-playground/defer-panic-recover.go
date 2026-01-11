package main

import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker(){
	fmt.Println("About to panic")
	defer func ()  {
		if err := recover();err != nil{
			log.Println("Error:",err)
		}
	}()
	panic("Something bad happened !")
}