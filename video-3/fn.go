package main

import (
	"errors"
	"fmt"
)

type argFNInt func(i, j int) (string, error)

func main() {
	s, err := fh(func(i, j int) (string, error) {
		if i > j {
			return "i is greater", nil
		}

		return "", errors.New("j cannot be greater")
	})

	fmt.Printf("%v, %v", s, err)
}

func fh(a argFNInt) (string, error) {
	s, err := a(99, 2)

	if err != nil {
		return "", errors.New("something wrong")
	}

	return s, nil
}
