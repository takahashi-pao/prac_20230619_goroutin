package main

import (
	"fmt"
	"time"
)

func main() {
	go count("goroutine")
	count("main")

	time.Sleep(time.Second)
}

func count(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(label, i)
		time.Sleep(time.Microsecond * 500)
	}
}
