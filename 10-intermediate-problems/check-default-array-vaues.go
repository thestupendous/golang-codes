package main

import "fmt"

func main() {
	var a [10]int
	for _, val := range a {
		fmt.Printf("%v ",val)
	}
}
