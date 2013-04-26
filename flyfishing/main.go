package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.Lake{MaxLocation: flyfishing.Location{600, 400}}
	lake.FillWithFishes(100)
	log.Println(len(lake.Fishes))
}