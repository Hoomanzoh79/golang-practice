package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened !")
}
