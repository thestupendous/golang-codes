package main

import "fmt"

func main(){
   for i:=0; i<34; i=i+5{     //traditional for loop
      fmt.Println(i)
   }

      fmt.Println("-*-8_80-8-8-8-8---8--8-8-8--8-8--88-")

   i:=30
   for ; i>0 ; i-=5{       //tweaking for to work as while
      fmt.Println(i)
   }

   fmt.Println("-*-8_80-8-8-8-8---8--8-8-8--8-8--88-")

   sum:=1
   for sum<1025{          //while loop
      fmt.Println(sum)
      sum+=sum
   }

   var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

   fmt.Println("now range loop")
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}