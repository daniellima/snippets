package main

import (
	"errors"
	"fmt"

	goerrors "github.com/go-errors/errors"
)

func ErrorHandlingMain() {
	err := fmt.Errorf("just a simple error")
	err2 := fmt.Errorf("i just show this error: %v", err)
	err3 := fmt.Errorf("i wrap this error: %w", err)

	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Printf("Error 3 wraps this: %v\n", errors.Unwrap(err3))

	gerr := goerrors.Errorf("just a error with StackTrace")
	gerr2 := goerrors.Errorf("i just show this error: %v", gerr)
	gerr3 := goerrors.Errorf("i wrap this error: %w", gerr)

	fmt.Println(gerr.ErrorStack())
	fmt.Println(gerr2.ErrorStack())
	fmt.Println(gerr3)
	fmt.Printf("Error 3 wraps this: %v\n", errors.Unwrap(gerr3))
	if errors.Is(gerr3, gerr2) {
		fmt.Println("gerr3 is a type of gerr2")
	} else {
		fmt.Println("gerr3 is not a type of gerr2")
	}
}
