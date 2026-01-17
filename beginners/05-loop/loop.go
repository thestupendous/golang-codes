package main
import "fmt"

func main() {
	for i:=0; i<5; i++ {
		fmt.Print(">")
	}
	fmt.Println()


	i:=5
	for i>0 {
		fmt.Print(">")
		i-=1
	}
	fmt.Println()
}
