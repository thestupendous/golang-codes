package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	b := 0
	a := 1
	return func() int {
        fmt.Println("a , b ",a,b)
		a,b = a+b, a
//        if a==b{
 //           return 0
  //      }
		return a
}
}

func main() {
	f := fibonacci()
    fmt.Println("printing first 10 items")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

