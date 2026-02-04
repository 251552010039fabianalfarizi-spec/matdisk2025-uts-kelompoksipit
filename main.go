package main

import (
	"fmt"
)

func main() {
	graph := Graph{
		"Hub Depok": {
			{"Citayam", 7},
			{"Bojong Gede", 10},
		},
		"Citayam": {
			{"Hub Depok", 7},
			{"Hub Cibinong", 6},
		},
		"Bojong Gede": {
			{"Hub Depok", 10},
			{"Hub Cibinong", 8},
		},
		"Hub Cibinong": {
			{"Citayam", 6},
			{"Bojong Gede", 8},
		},
	}

	start := "Hub Depok"
	end := "Hub Cibinong"

	dist := Dijkstra(graph, start)

	fmt.Println("=== Optimasi Jalur Terpendek (Graph) ===")
	fmt.Println("Dari :", start)
	fmt.Println("Ke   :", end)
	fmt.Println("Jarak Terpendek:", dist[end], "km")
}
