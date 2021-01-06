package main

import (
	"log"

	"github.com/lc/subjs/runner/subjs"
)

func main() {
	opts := subjs.ParseOptions()
	runner := subjs.New(opts)
	err := runner.Run()
	if err != nil {
		log.Fatalf("Error running subjs: %s", err)
	}
}
