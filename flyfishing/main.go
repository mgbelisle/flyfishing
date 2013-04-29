package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.NewLake()
	casts := castNTimesAsync(100, lake)
	log.Println(len(casts))
}

type cast struct {
	loc flyfishing.Location
	fish flyfishing.Fish
}

func castNTimes(n int, lake flyfishing.Lake) []cast {
	casts := []cast{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		casts = append(casts, cast{loc, fish})
	}
	return casts
}

func castNTimesAsync(n int, lake flyfishing.Lake) []cast {
	castChan := make(chan cast)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
			castChan <- cast{loc, fish}
		}()
	}
	casts := []cast{}
	for i := 0; i < n; i++ {
		casts = append(casts, <-castChan)
	}
	return casts
}
