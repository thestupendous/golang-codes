/*
	use of functions in golang
	1. return types - golang supports multiple return types for a function
	2. taking pointers as function parameters
	3. variadic functions - those which accept unknown number of args (like fmt.Println)
	4. anonymous functions in golang
*/

package main
import "fmt"
func swap(a,b int) (int,int) {			//multiple return values
	return b,a}

func swap_using_pointers(a,b *int) {	//function taking pointers as args
	t:= *a
	*a = *b
	*b = t
}

func sum(a ...int) int {				//function accepting indefinite number of arguments - variadic function
	s := 0
	for _ , marks := range a {
		s += marks
	}

	//fmt.Printf("type of sum is %T\n",s)
	return s
}

func main(){
	defer fmt.Println("bye bye")

	var line func() = func() { fmt.Println() }			//anonymous function defined and assigned to line variable
	a:=3
	b:=1000

	a,b = swap(a,b)
	fmt.Println("after first swap, a , b = ",a,b)
	line()

	swap_using_pointers(&a,&b)
	fmt.Println("after second swap, a , b = ",a,b)
	line()
	line()


	//variadic func
	arith_marks := []int {42,44,25,30,49,50,23}
	//fmt.Println("sum ",sum(arith_marks...))
	fmt.Println("avg marks of class A in Arithmetics are ",( float64(sum(arith_marks...))/ float64(len(arith_marks)) ) )
	line()
	fmt.Println("sum of 9 93 9 2 0 -34 is ",sum(9,93,9,2,0,-34))
	line()
}

