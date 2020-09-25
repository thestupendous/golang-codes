package main


import "fmt"

func main(){
    for i:=0 ; i<10 ; i=i+1 {
        for j:=0 ; j<=i ; j=j+1 {
            fmt.Printf("*")
        }
        fmt.Printf("\n")
   }} 
