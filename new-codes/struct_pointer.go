package main

import "fmt"

type mod struct {
	name string
	version uint32
}

func main() {
	invis := mod{"Invisible",4}
	god := &mod{"God abilities",4}

	fmt.Printf("%T %v \n%v\n",god,god,invis)
}
