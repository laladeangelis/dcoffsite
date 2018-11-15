package main

import (
	"fmt"
)

func helloDC(done chan bool) {
	fmt.Println("Tom is Awesome!")
	done <- true
}
func main() {
	done := make(chan bool)
	go helloDC(done)
	<-done
	fmt.Println("main function")
}
