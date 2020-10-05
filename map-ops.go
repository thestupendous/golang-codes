package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42			//retrieve element
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48			//retrieve element
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")			//delete  element
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]			//check presence of a key
	fmt.Println("The value:", v, "Present?", ok)
}
