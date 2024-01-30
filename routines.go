package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("process")

	go f("processOtherThread1")

	go func(msg string) {
		fmt.Println(msg, ":", 1)
	}("processOtherThread2")
	time.Sleep(time.Second)
	fmt.Println("done")
}
