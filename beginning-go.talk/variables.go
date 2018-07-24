package main

import "fmt"

const TotalBottles = 100

func main() {
	var fallen = 0
	fmt.Println(TotalBottles-fallen, "bottles on the wall")
	fallen = fallen + 1
	fmt.Println(TotalBottles-fallen, "bottles on the wall")
}
