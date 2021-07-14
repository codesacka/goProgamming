package main

import "errors"

type error interface {
	Error() string
}

func foo() error {
	return errors.New("Some Error Ocurred")
}

func main() {

	if err := foo(); err != nil {
		//Handle the error
	}

}
