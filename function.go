package main

import "fmt"


func main(){
    var minFunc func(int,int) int      //using a function as a variable
    minFunc = func(a int, b int) int { //using function as a value
        return 2*a + b-a 
    }
    fmt.Println(minFunc(3,9))
    
    func (msg string){
        fmt.Println("the message is ",msg)
    } ("cyan is Imposter")

    fmt.Println("done")

}
