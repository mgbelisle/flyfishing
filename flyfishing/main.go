package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.Lake{MaxLocation: flyfishing.Location{600, 400}}
	for i := 0; i < 100; i++ {
		lake.AddFish()
	}
	log.Println(len(lake.Fishes))
}