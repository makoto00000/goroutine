package main

import (
	"fmt"
	"log"
	"time"
)

func timeout(second int) {
	log.Print(fmt.Sprintf("sleep%d started.", second))
	time.Sleep(time.Duration(second) * time.Second)
	log.Print(fmt.Sprintf("sleep%d finished", second))
}

func main() {

	log.Print("started.")
	start := time.Now()

	finished := make(chan bool)

	funcs := []func(){
		func() {
			timeout(1)
			finished <- true
		},
		func() {
			timeout(2)
			finished <- true
		},
		func() {
			timeout(3)
			finished <- true
		},
	}

	for _, sleep := range funcs {
		go sleep()
	}

	for i := 1; i <= len(funcs); i++ {
		<- finished
	}

	log.Print("all finished.")
	end := time.Now()

	log.Printf("%.0f seconds elapsed.", (end.Sub(start).Seconds()))

}