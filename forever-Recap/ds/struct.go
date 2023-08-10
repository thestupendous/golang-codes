package main

import "fmt"

type bhagona struct {
	radius float32
	height float32
}

func bada(a bhagona,b bhagona) bool {
	//comparing by volume
	
	if ( a.radius * a.radius * a.height > b.radius * b.radius * b.height  ) {
	  return true }
	return false
//	return true

}

func incHeight(a *bhagona) {
	(*a).height += 1
}




func main() {
	fmt.Println(bhagona{3.4,2.5})
	bhagona2 := bhagona{}
	var bhagona3 bhagona

	bhagona2.radius = 5
	bhagona2.height = 3.7
	
	fmt.Println(bhagona2)
	fmt.Println(bhagona3)

	res := bada(bhagona2,bhagona3)
	fmt.Println("is bhagona2 > bhagona3 ",res)
	
	fmt.Println("before increment bhagona2 ",bhagona2)
	incHeight(&bhagona2)
	fmt.Println("after incrementing bhagona2 ",bhagona2)
}


