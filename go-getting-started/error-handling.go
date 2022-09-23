package main

import (
	"errors"
	"fmt"
)

func ErrorHandlingMain() {
	err := fmt.Errorf("just a simple error")
	err2 := fmt.Errorf("i just show this error: %v", err)
	err3 := fmt.Errorf("i wrap this error: %w", err)

	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Printf("Error 3 wraps this: %v\n", errors.Unwrap(err3))
}
