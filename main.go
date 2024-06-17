package main

import (
	"log"
	"oddshub/ui"
)

func main() {

	if err := ui.RunApp(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
