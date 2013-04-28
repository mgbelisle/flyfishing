package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.Lake{MaxLocation: flyfishing.Location{600, 400}}
	for i := 0; i < 100; i++ {
		lake.Fishes = append(lake.Fishes, flyfishing.NewCutthroat(lake))
	}
	log.Println(len(lake.Fishes))
}