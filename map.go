package main
import "fmt"

type person struct {
	 name string
	 age int
}

var m map[string]person
func main(){
	m = make(map[string]person)
	m["good person"] = person{"anand",22}
	m["beautiful"] = person{"sweety",22}

	fmt.Println(m["beautiful"])
	fmt.Println(m["good person"])
}