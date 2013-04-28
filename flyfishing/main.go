package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.NewLake()
	for i := 0; i < 100; i++ {
		fish := lake.Cast(flyfishing.Caddis{}, lake.RandLoc())
		if fish != nil {
			log.Println(fish)
		}
	}
}