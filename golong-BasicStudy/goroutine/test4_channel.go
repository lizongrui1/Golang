package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	x, y
}
