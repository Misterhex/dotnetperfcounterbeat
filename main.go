package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/mavlyutov/golagraphite"
)

func main() {

	counterName := `\.NET CLR Memory(*)\# Gen 0 Collections`

	c, err := golagraphite.ReadPerformanceCounter(counterName, 10)

	if err != nil {
		log.Printf("unable to read perf counter \n")
	}

	for m := range c {
		spew.Dump(m)
	}
}
