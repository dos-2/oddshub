package main

import (
	"log"

	"github.com/dos-2/oddshub/ui"
)

func main() {

	if err := ui.RunApp(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
