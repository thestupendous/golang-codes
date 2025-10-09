/*
	using struct for storing sign and magnitude of any integer number
	using uint32 for storing magnitude of numbers
*/

package main

import "fmt"

type num struct{
	sign bool
	mag uint32
}
func (n num) print() {
	numstr := ""
	if !bool {
		numstr += "-"
	}
	numstr+= string(mag)
	fmt.Println(numstr)
}

func add(a,b num) num {
	sig := false
	ma := uint32(0)

	if a.sign && b.sign {
		ma = a.mag+b.mag
		sig = a.sign
	} else {
		ma = a.mag-b.mag
		if ma<0 {
			ma = -ma
		}
		if a.mag>b.mag {
			sig = a.sign
		}else{
			sig = b.sign
		}
	}
	return num{sig,ma}
}

func (a num) mul(b num) num {										//use of methods
	res := num{sign: false, mag: 0}
	res.mag = a.mag * b.mag
	res.sign = a.sign && b.sign

	return res
}

func (a num) div(b num) num{										//use of methods
	res:= num{ sign: (a.sign && b.sign) , mag: (a.mag/b.mag) }

	return res
}

func sub(a,b num) num {

	b.sign = !b.sign
	return add(a,b)

}


func main() {

	var a num = num{true, 2300}
	var b num = num{false, 300}
	c := num{false, 30}
	_ = c

//	fmt.Println(a)
	a.print()
	fmt.Println(b)
	fmt.Println("sum of a and b ",add(a,b))
	fmt.Println("difference of a and b ",sub(a,b))
	fmt.Println("product of a and b ",a.mul(b))
	fmt.Println("division of a and b ",a.div(b))

}
