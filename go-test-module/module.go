package module

import (
	"errors"
	"fmt"
)

func SayHelloTo(name string) (string, error) {

	if name == "" {
		return "", errors.New("you must say who you are saying hello to")
	}

	return fmt.Sprintf("Hello, %s!!!", name), nil
}

func SayTheQuote(quoteNumber int) string {
	quotes := []string{
		"1+1 = 2",
		"Tudo se transform",
		"Em cima, é como embaixo",
		"Go is a nice language",
	}

	return quotes[quoteNumber%len(quotes)]
}
