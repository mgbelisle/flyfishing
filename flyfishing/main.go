package main

import (
	"flyfishing"
	"log"
)

func main() {
	lake := flyfishing.NewLake()
	castResults := castNTimes(100, lake)
	log.Println(len(castResults))
}

type castResult struct {
	loc flyfishing.Location
	fish flyfishing.Fish
}

func castNTimes(n int, lake flyfishing.Lake) []castResult {
	castResults := []castResult{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		castResults = append(castResults, castResult{loc, fish})
	}
	return castResults
}

func castNTimesAsync(n int, lake flyfishing.Lake) []castResult {
	castResults := []castResult{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		castResults = append(castResults, castResult{loc, fish})
	}
	return castResults
}
