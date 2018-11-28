package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("bum")
			break
		}
		println(timer)
		time.Sleep(1 * time.Second)

	}
}
