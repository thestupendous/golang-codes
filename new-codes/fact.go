package main

import "fmt"

func fact(n uint64) uint64 {
	if n<=1 {
		return 1
	}

	return n*fact(n-1)

}

func main() {

	for i:=1; i<=10; i++ {
		fmt.Println(i," ",fact(uint64(i)))
	}

}
