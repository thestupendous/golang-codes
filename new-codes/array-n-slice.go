package main
import "fmt"

func main(){
	var a [4]int			//declared an array
	a[0]=1
	fmt.Printf("%T\n",a)
	//append(a,90)
	fmt.Println(a[0:2])
	s := a[0:2]				//declared slice
	fmt.Printf("%T\n",s)


	var s2 []int			//another slice declaration
	s2 = append(s2,9,2,-23,88)
	fmt.Printf("%T\t%v",s2,s2)

}
