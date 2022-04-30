package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		defer fmt.Println("Timer 1 fires(goroutine1 finished)")
		<-timer1.C
	}()

	timer2 := time.NewTimer(time.Second)
	go func() {
		defer fmt.Println("Timer 2 fires(goroutine2 finished")
		<-timer2.C
	}()
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
	time.Sleep(2 * time.Second)
}
