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

	sleep1_finished := make(chan bool)
	sleep2_finished := make(chan bool)
	sleep3_finished := make(chan bool)

	go func() {
		timeout(1)
		sleep1_finished <- true
	}()
	go func() {
		timeout(2)
		sleep2_finished <- true
	}()
	go func() {
		timeout(3)
		sleep3_finished <- true
	}()

	<- sleep1_finished
	<- sleep2_finished
	<- sleep3_finished

	log.Print("all finished.")
	end := time.Now()

	log.Printf("%.0f seconds elapsed.", (end.Sub(start).Seconds()))

}