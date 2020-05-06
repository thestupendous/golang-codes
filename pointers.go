package main

import "fmt"

func main(){
	var a,b,c int
	p := &a
	a = 3
	b = 4
	c = 5
	fmt.Println(*p)
	*p = 90
	fmt.Println(a)
	a=b+c
}