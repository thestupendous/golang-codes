package main

import "fmt"

func main(){
	var st string
	st = "कanand kumar  "

	for _,ch := range st {
		fmt.Print(string(ch)," ")
	}
	fmt.Println()

	unknown := st[0]
	fmt.Printf("%T\n",unknown)
	fmt.Println(len(st))
	fmt.Println(string(st[:4]))



}
