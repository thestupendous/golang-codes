//this program counts the frequency of every word in a string
package main

import (

	// "golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	rist := strings.Fields(s)
	fmt.Printf("rist %v , and type %T , done\n",rist,rist)
	for _ , i := range rist{
		m[i] += 1
	}
	
	return m
	
	
}

func main() {
	fmt.Println(WordCount("abc bca dbc adf abc"))
	// wc.Test(WordCount)
}
