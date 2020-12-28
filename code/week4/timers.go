package main

import (
	"fmt"
	"time"
)

func main() {
	firstTimer := time.NewTimer(2 * time.Second)

	<-firstTimer.C

	fmt.Println("Timer 1 fired")

	secondTimer := time.NewTimer(time.Second)

	go func() {
		<-secondTimer.C
		fmt.Println("Timer 2 fired")
	}()

	stopSecond := secondTimer.Stop()

	if stopSecond {
		fmt.Println("Timer 2 stoped...")
	}

	time.Sleep(time.Second * 2)

}
