package main

import (
	"log"
	"oddshub/ui"

//	"fmt"
	"oddshub/API"
	"oddshub/models"
//	"oddshub/slides"
//	"oddshub/sports"
)

func main() {
	// Load sports events
//	events := loadEvents()
//	games := events[string(sports.Americanfootball_nfl)]
//	var result string
//	result += slides.FormatTeamEvent(games[0])
//	fmt.Print(result)
 // Run the presentation

	if err := ui.RunApp(); err != nil {
	  log.Fatalf("Failed to run application: %v", err)
	}
}

func loadEvents() map[string][]models.Event {
	events := API.GetAllUpcomingEventsMap()
	if events == nil {
		panic("Failed to load events")
	}
	return events
}
