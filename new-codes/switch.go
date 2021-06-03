package main

import (
	"fmt"
	 "time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday , time.Sunday :
		fmt.Println("it's a weekend today, enjoy and take a break!")
		break
		
	default:
		fmt.Println("it's a weekday, have a great day!")

	}


}
