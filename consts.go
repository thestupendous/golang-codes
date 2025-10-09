package main

import (
	"fmt"
	"math"
)


func main(){
	const size = 33
	fmt.Println(size)

	const big = 8.92302e12
	fmt.Printf("%T is type, %v is value\n",big,big)

	fmt.Println(int64(big))

	fmt.Println(math.Abs(math.Sin(890)))


}
