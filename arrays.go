package main 
import (
	"fmt"
)
const SIZE=4

func main(){

	var yname string
	fmt.Print("enter your name ")
	fmt.Scanln(&yname)
	fmt.Print("hello "+yname+"! Now please enter your ",SIZE," bff names ")
	var names [SIZE] string
	for i:=0;i<SIZE;i++{
		fmt.Scan(&names[i])
	}

	for i:=0;i<SIZE;i++{
		fmt.Println(names[i])
	}


}
