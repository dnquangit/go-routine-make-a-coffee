package main

import (
	"go-module/coffee"
	"time"
)

func main() {
	go coffee.MakeASimpleCupOfCoffee()

	/* Wait a cup of coffee */
	time.Sleep(time.Second * 150)
}
