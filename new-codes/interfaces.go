package main

import "fmt"
import "strconv"

type array_ops interface {
	sum() int
	mean() float64
	mod() int
}

type myarr struct {
	l int
	arr []int
}

func (ma myarr) String() string {				//stringer interface used to return string of struct
	var out string
	out = "[ "
	for _,i := range(ma.arr) {
		out += strconv.Itoa(i) + " "
	}
	out += "]"
	return out
}


func (a myarr) sum() int {
	var som int =0
	for _,num := range a.arr{
		som += num
	}
	return som
}

func (a myarr) mean() float64{
	var som int = a.sum()
	var res float64
	res = float64(float64(som)/float64(len(a.arr)))

	return res
}

func (a myarr) mod() int {
	return 0
}

func arrView(aobj array_ops) {
	fmt.Println("for current array ")
	fmt.Println("sum = ",aobj.sum())
	fmt.Println("mean = ",aobj.mean())
}

func main() {
	arr1 := myarr{ arr: []int{3,4,2,5,1,6} }
	arr1.l = len(arr1.arr)
	fmt.Println(arr1)
	fmt.Println(arr1.sum())
	fmt.Println(arr1.mean())

	var aa array_ops
	_ = aa

	arrView(arr1)

//	var nops array_ops = 
}

