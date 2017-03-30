package main

import "time"
import "fmt"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	now := <-timer1.C
	now1 := time.Now()
	fmt.Printf("now: %v vs %v\n", now, now1)
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
