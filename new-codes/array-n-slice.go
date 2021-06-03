package main
import "fmt"

func main(){
	var a [4]int
	a[0]=1
	fmt.Printf("%T\n",a)
	//append(a,90)
	fmt.Println(a[0:2])
	s := a[0:2]
	fmt.Printf("%T\n",s)


	var s2 []int
	s2 = append(s2,9,2,-23,88)
	fmt.Printf("%T\n",s2)

}
