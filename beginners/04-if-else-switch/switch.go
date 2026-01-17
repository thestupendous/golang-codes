package main
import "fmt"

func main() {

	n := 3
	result := ""

	switch n {
	case 1:
		result = "one"
	case 2:
		result = "two"
	case 3:
		result = "three"
	case 4:
		result = "four"
	case 5,6,7:
		result = "char se jada"
	default:
		result = "NAN"
	}

	fmt.Println(result)

}
