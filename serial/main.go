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

	timeout(1)
	timeout(2)
	timeout(3)

	log.Print("all finished.")
	end := time.Now()

	log.Printf("%.0f seconds elapsed.", (end.Sub(start).Seconds()))

}