package main

import (
	"runtime"
	"time" s
)

func main() {
	runtime.GOMAXPROCS(8)

	ch := make(chan string)

	go abcGen(ch)
	go printer(ch)

	println("bekleme")
	time.Sleep(100 * time.Microsecond)
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string) {
	for {
		println(<-ch)
	}

}
