package main

import (
	"errors"
	"fmt"
)

var FailedPaymentError = &paymentError{msg: "your cc is not available"}
var InsFoundsError = &paymentError{msg: "your founds are 0 hahaha"}

type paymentError struct {
	msg string
}

func (e *paymentError) Error() string {
	return e.msg
}

func main() {
	//
	//
	_, err := stub(1)

	/*if err != nil {
		if errors.Is(err, FailedPaymentError) {
			fmt.Println("contact your bank please")
		}
		if errors.Is(err, InsFoundsError) {
			fmt.Println("you do not have any money")
		}
	}*/

	customerPaymentError := fmt.Errorf("the CC is expired %w", err)

	fmt.Println(errors.Unwrap(customerPaymentError))

}

func stub(i int) (int, error) {
	if i == 1 {
		return 1, InsFoundsError
	}
	return 0, FailedPaymentError
}
