package main

func main() {
	m := make(map[string]Vertex)
	m["Tehran"] = Vertex{
		Lat:  35.6892,
		Long: 51.3890,
	}
	printMap(m)
}
