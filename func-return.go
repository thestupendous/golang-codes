package main

import "fmt"

func swap(x,y string) (string, string) {	//this function returns two values
	return y,x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


func main(){
	var a,b string
	a= "hii"
	b= "there"
	a,b = swap(a,b)
	fmt.Println(b,a)
	fmt.Println(split(98))
}
