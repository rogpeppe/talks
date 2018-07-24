package main

import "fmt"

const TotalBottles = 100

func main() {
	var fallen int
	for fallen = 0; TotalBottles-fallen > 0; fallen++ {
		fmt.Println(TotalBottles-fallen, "bottles on the wall")
	}
	fmt.Println("no bottles on the wall")
}
