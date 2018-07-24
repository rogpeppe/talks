package main

import "fmt"

const TotalBottles = 100

func main() {
	var fallen = 0
	for TotalBottles-fallen > 0 {
		fmt.Println(TotalBottles-fallen, "bottles on the wall")
		fallen++
	}
	fmt.Println("no bottles on the wall")
}
