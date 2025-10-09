package main

import "fmt"
import "time"
var a int
func waitnReturn()  {
	time.Sleep(time.Millisecond*1000*2)
	fmt.Println("\t\troutine work finished, returning to main")
	a=5
}

func main() {
	go waitnReturn()
	a=0

	if a==5 {
		fmt.Println("found finished goroutine")
	} else {
		fmt.Println("still waiting in main for routine to finish :(")
	}
	time.Sleep(time.Millisecond*1000*1)

	if a==5 {
		fmt.Println("found finished goroutine")
	} else {
		fmt.Println("still waiting in main for routine to finish :(")
	}
	time.Sleep(time.Millisecond*1000*1)

	if a==5 {
		fmt.Println("found finished goroutine")
	} else {
		fmt.Println("still waiting in main for routine to finish :(")
	}
	fmt.Println("exiting main")
}



