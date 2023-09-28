package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	send := func(c <-chan interface{}) {
		for {
			select {
			case _, ok := <-c:
				if !ok {
					close(out)
					return
				}
			case _, ok := <-out:
				if !ok {
					return
				}
			}
		}
	}
	for _, c := range channels {
		go send(c)
	}
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))

}
