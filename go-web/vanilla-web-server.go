package main

import (
	"fmt"
	"net/http"
)

func VanillaWebServerMain() {
	fmt.Println("Hello, Go for the Web")

	http.HandleFunc("/test", func(response http.ResponseWriter, r *http.Request) {
		response.Write([]byte("Hello?"))
	})

	fmt.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf(err.Error())
}
