package main

import (
	"fmt"
	_ "strconv"
)
var omnipotent int = 90000

func main(){
	fmt.Println("go" + " through");

	fmt.Println("1+5 = ",1+5);

	fmt.Print("no next line \n");


	fmt.Println("7.0/3.0 = ",7.0/3.0);
	var number_string string
	temp := 7.0/3.0
	//number_string = strconv.Itoa(-89909)   //Itoa used for ints
	number_string = fmt.Sprintf("%f",temp)


	fmt.Println(number_string)
	fmt.Println(true && false)
	fmt.Println(!true)


	var g = 980
	fmt.Println(g)
	var h,i int = 892, -34
	fmt.Println(h,i)
	var  name = "Rajwati"
	var  age uint64
	avogadro_number := 3.5e20
	spec := "auto"
	_ = name ; _ = age ; _ = spec		//to get rid of unused vars warning
	fmt.Println(avogadro_number)

}
