package main

import (
	"fmt"
)

func main(){
	var a [2] string
	a[0] = "hii"
	a[1] = "there"

	fmt.Println(a)

	var primes [6] int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7
	primes[4] = 11
	primes[5] = 13

	fmt.Println(primes)

	days := [7] string {"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	fmt.Println(days)

	fmt.Println("from tuesday to saturday",days[1:6])				//slice


}