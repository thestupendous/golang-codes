package main

import (
	"fmt"
)

func add(x int, y int) int{
	return x+y
}

func subtract(x int, y int) int{
	return x-y
}

func multiply(x int, y int) int{
	return x*y
}

func divide(x int, y int) int{
	return x/y
}
func main(){
	var x ,y int
	x = 5
	y  = 2
	fmt.Println("numbers ",x," ",y)
	fmt.Println("addition ",add(x,y))
	fmt.Println("subtraction ",subtract(x,y))
	fmt.Println("multiplication ",multiply(x,y))
	fmt.Println("division ",divide(x,y))
}