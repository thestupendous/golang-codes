package main
import "fmt"

func fib(n int ) int {
	if (n==1 || n==2) {
		return 1
	}
	n-=2
	n1:=1
	n2:=1
	for ; n>0; n--{
		t:= n2
		n2 += n1
		n1 = t
	}
	return n2
}
func main() {
	fmt.Println(fib(5))
}
