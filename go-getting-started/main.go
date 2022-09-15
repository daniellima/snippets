package main

import (
	"fmt"
	"log"

	module "daniellima.dev/go-test-module"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println(quote.Glass())

	gretting, err := module.SayHelloTo("Wildlife")
	if err != nil {
		log.Fatal("Error when gretting Wildlife")
	}

	quotesIwant := []int{0, 1, 3, 2, 1}
	for i, quoteNumber := range quotesIwant {
		fmt.Printf("Quote %v", i)
		fmt.Println(" -> ", module.SayTheQuote(quoteNumber))
	}
	fmt.Println(gretting)
}
