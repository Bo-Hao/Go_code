package main

import (
	"fmt"
	"time"
)

//goroutine is like thread

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}

}

func main() {
	go f(0)
	time.Sleep(time.Second * 1) //pause one second
}
