package main

import (
	"fmt"
	"time"
)
func doPrint(){
	for {
	fmt.Println("\033[H\033[2J")
	fmt.Println("Hello, welcome to the game this is the board []")
	time.Sleep(time.Millisecond*500)
}
}
func main() {
	i :=1
	j :=0
	go doPrint()

	for {
		fmt.Scanln(&j)
		i+=1
		if i%j==0 {
			fmt.Println("j is divisor of i!")
		}else {
			fmt.Println("NOOO, j is not a divisor of i")
		}
		if i%5==0 {
			fmt.Println("i is a new multiple of 5")
		}
		time.Sleep(time.Millisecond*300)
	}

}

