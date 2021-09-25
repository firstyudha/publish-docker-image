package main

import (
	"fmt"
	"publish-docker-image/calculator"
)

func main() {
	fmt.Println("Hello!")
	fmt.Println(calculator.Add(5, 3))
}
