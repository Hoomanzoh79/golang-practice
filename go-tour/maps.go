package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func printMap(m map[string]Vertex) {
	fmt.Println(m)  
}
