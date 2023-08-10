/*
	slices are like mutable data objects in python, meaning they always copy by reference,
	meaing, changing elements of one copy change them in other copies

	a slice's right sentinel can be extended or contracted any time
	    ex. []a{1,2,3,4}
		if you did a=a[:2]  (now a {1,2})
		 then did a=a[:4]  (now a {1,2,3,4})
		  It would still work
	but left sentinel if moved to right, stays there forever
		ex. []a{1,2,3,4}
		if you did a=a[1:]
	     it will never include the original a[0] i.e. '1' value
	note: these examples have been tested in important point 1, below

*/

package main

import "fmt"

func main(){
	a := [4]int{2,3,4,5}
	fmt.Println(a)

	b := a[1:2]
	fmt.Println(b[0])
	b[0]=-23
	fmt.Println(a[1])


	fmt.Printf("%T %T \n",a,b)
	fmt.Println(a)
	

	//important point 1
	//a slice's right sentinel can be extended or contracted any time
	b=b[0:0]
	fmt.Println(b)
	b=b[0:3]
	fmt.Println("capacity ",cap(b),"  b : ",b)
	b=b[:2]
	fmt.Println(b)

}



