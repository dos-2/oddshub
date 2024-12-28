/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package main

import (
	"log"
	"os"

	"github.com/dos-2/oddshub/models"
	"github.com/dos-2/oddshub/ui"
)

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "debug":
				models.SetDebug(true)
			}
		}
	}

	if err := ui.RunApp(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
