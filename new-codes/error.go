package main

import (
	"fmt"
	"errors"
)

func makeError() (int,error) {
	return 9, errors.New("created my own error")
}

func makeNoError() (int,error) {
	return 9, nil
}

func divide(a,b int) (int,error) {
	if b==0 {
		errCode := 9
		return -9999 , fmt.Errorf("ERRCODE %v (Division by zero) ",errCode)
	}

	return a/b, nil
}

func main() {

	var val int ; var err error

	val,err = makeError()
	fmt.Println(val,err)
	val,err = makeNoError()
	fmt.Println(val,err)

	val,err = divide(3,0)
	if err!=nil {
		fmt.Println("Error while division : ",err)
	} else {
		fmt.Println("result of division = ",val)
	}


}
