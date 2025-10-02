package main

import  (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	var map1 map [string]int  = make (map[string]int)

	inputString = "a b c d c b e f a c f e k c b g"
	var wordList []string = strings.Fields(inputString)

	for _,word := range wordList {
		map1[string(word)] += 1
	}

	for word, count := range map1 {
		fmt.Println(word,"came",count,"times")
	}


	_ = map1
}
