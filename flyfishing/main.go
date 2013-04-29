package main

import (
	"flyfishing"
	"log"
)

type cast struct {
	loc flyfishing.Location
	fish flyfishing.Fish
}

func main() {
	lake := flyfishing.NewLake()
	casts := []cast{}
	for i := 0; i < 100; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, loc)
		casts = append(casts, cast{loc, fish})
	}
	log.Println(casts[:10])
}