package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func main() {
	err := terror()

	if err != nil {
		fmt.Println("time is not even")
	}

	res, err := division(10, 2)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

	_, err = division(10, 0)

	if _, errr := division(10, 0); errr != nil {
		fmt.Println("cannot show result")
	}

}

// the interfaces in golang are pointers
func terror() error { // error is an interface
	if time.Now().Unix()%2 == 0 {
		return nil
	}
	return &MyError{Msg: "this is bad"}
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func division(a, b int) (float32, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}

	return float32(a) / float32(b), nil
}
