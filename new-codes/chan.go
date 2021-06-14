package main

import "fmt"
import "time"

func work(message chan string) {
	i := 0
	message <- fmt.Sprint("i is ",i)
	time.Sleep(time.Second)
	i+=1
	message <- fmt.Sprint("i is ",i)
	time.Sleep(time.Second)
	i+=1
	message <- fmt.Sprint("i is ",i)
	close(message)
}


func main() {
	fmt.Println("hi")

    message := make(chan string)

	go work(message)

	for {
	    msg, open := <-message
		if !open { break }
		fmt.Println(msg)
	}
	fmt.Println("bye")

}
