package main
import "fmt"

func add(a int,b int) int {
	return a+b
}
func multiReturns(a int,b int) (int, int, int) {
	return a+b, 0, 1
}

func main() {
	result := add(1,4)
  fmt.Println(result)
}
