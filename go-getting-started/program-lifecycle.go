package main

import (
	"fmt"
	"time"
)

func IntProducer() chan int {
	ch := make(chan int)
	c := 0
	go func() {
		for {
			c++
			ch <- c
		}
	}()
	return ch
}

func PackageLifeCycleMain() {

	c := make(chan int)

	// go func() {
	// 	for {
	// 		fmt.Println("starting from inside the goroutine")
	// 		value, ok := <-c
	// 		time.Sleep(10 * time.Millisecond)
	// 		fmt.Printf("finishing from inside the goroutine. Value: %v, Channel closed? %v\n", value, !ok)
	// 	}
	// }()
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	fmt.Println("oi")

	c <- 5
	c <- 50
	c <- 70
	close(c)

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("oi?")

	fmt.Println("oi+")

	producer := IntProducer()

	for i := 0; i < 3; i++ {
		fmt.Println(<-producer)
	}
}
