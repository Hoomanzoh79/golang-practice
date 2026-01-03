package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func PrintMap(m map[string]Vertex) {
	fmt.Println(m)  
}