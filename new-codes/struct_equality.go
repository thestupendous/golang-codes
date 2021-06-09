package main

import "fmt"
import "reflect"

type person struct{
	name string
}

func main(){

	p1 := person{"क म ल"}
	p2 := person{"के र ल"}
	p3 := person{"क म ल"}

	fmt.Println("p1 , p2, p3 ",p1," ",p2," ",p3)
	fmt.Println("is p1 == p3 , using operator ",p1==p3)
	fmt.Println("is p1 == p3 , using DeepEqual() ",reflect.DeepEqual(p1,p3))
}
