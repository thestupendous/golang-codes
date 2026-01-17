package main
import "fmt"
func main() {
	type UserID int
	type Direction byte
	var p1 UserID = 123
	var left Direction = 1
	fmt.Println(p1,left)

	/*  type conversion */
	var p2 UserID = UserID(201)
	var right Direction = Direction(2)
	fmt.Println(p2,right)
}
