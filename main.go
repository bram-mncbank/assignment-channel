package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 1; i < 10; i++ {
		fmt.Printf("Player %v start\n", i)
		go bolaPanas(i, ch)
	}
	fmt.Printf("Player %v dead!", <-ch)
}

func bolaPanas(p int, c chan int) {
	for {
		time.Sleep(1 * time.Second)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		t := r.Intn(100) % 10
		fmt.Printf("Player %v hold the bom...\n", p)
		if t == 0 {
			c <- p
			return
		}
	}
}
