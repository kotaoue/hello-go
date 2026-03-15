package main

import (
	"log"

	"github.com/kotaoue/goimport/packages/car"
	"rsc.io/quote"
)

func main() {
	log.Println("Start")
	log.Println(quote.Hello())
	log.Println(car.Ignition())
	log.Println("End")
}
