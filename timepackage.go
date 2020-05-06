package main
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Tuesday)
	fmt.Println(time.Tuesday+2)
	// var a int
	// a = nil
	// fmt.Println(a)
}
