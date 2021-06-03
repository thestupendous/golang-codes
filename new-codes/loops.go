package main

import (
	 "fmt"
)
func line(){
	fmt.Println()
}

func main() {
	i:= 99;

	for i>0 {
		fmt.Print(i," ")
		i/=2
	}
	line()

	for i:=0; i<25; i+=1 {
		fmt.Print(i," ")
	}
	line()


	for i:=98; i>2; i-- {
		fmt.Print(i," ")
		if i%33 == 0 {
			break}
		}
	line()

	fmt.Print("printing evens ")
	for i:=1; i<=100; i++ {
		if i%2!=0 { continue}
		fmt.Print(i," ")
	}
	line()

}
