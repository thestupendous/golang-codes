package main

import (
	"fmt"
	"math/rand"
)

func main(){

	fmt.Println("random number for you ",rand.Intn(99999))
	for i:=0;i<10;i++ { fmt.Print(rand.Int(2)," ") }
}
