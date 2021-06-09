//bubble sort algorithm implementation in go
//sorting in decreasing order
package main
import "fmt"

func swap(a,b *int) {
	//t := *a ; *a = *b ; *b = t }		//swapping using temp variable
	*a , *b = *b, *a}					//swapping using multiple assignment feature of go (like python)


func main() {
	var a[] int
	a = append(a,2,3,4,5,6,4,10,90)			//defining a slice with static values (sorted in increasing order)
	l := len(a)
	fmt.Println(a)

	for i:=0 ; i<l-1; i++ {					//outer loop for finding out min value which bubbles to the end
		for j:=0; j<l-1-i ; j++ {			//inner loop performing swaps to find next min in each iteration of outer loop
			if a[j]<a[j+1] {
				swap(&a[j],&a[j+1])
			}
		}
	}


	fmt.Println(a)

}
