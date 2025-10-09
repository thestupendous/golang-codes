package main
import "fmt"

func fibo(n int) int{
	if n<=1 {
		return 1
	}else{
		return n + fibo(n-1)
	}
}

func main(){

	for i:=0 ; i<=8 ; i++ {
		fmt.Print(fibo(i)," ")
	}

}
