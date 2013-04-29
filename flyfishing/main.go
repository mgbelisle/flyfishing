package main

import (
	"flyfishing"
)

func main() {
	lake := flyfishing.NewLake()
	casts := castNTimesAsync(100, lake)
	lake.ShowCasts(casts)
}

func castNTimes(n int, lake flyfishing.Lake) []flyfishing.Cast {
	casts := []flyfishing.Cast{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		casts = append(casts, flyfishing.Cast{loc, fish})
	}
	return casts
}

func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.Cast {
	castChan := make(chan flyfishing.Cast)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
			castChan <- flyfishing.Cast{loc, fish}
		}()
	}
	casts := []flyfishing.Cast{}
	for i := 0; i < n; i++ {
		casts = append(casts, <-castChan)
	}
	return casts
}
