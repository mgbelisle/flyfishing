package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.NewLake()
	numCasts := 100
	for i := 0; i < numCasts; i++ {
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		if fish != nil {
			log.Println(fish)
		}
	}
}