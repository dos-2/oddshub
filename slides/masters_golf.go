/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"github.com/dos-2/oddshub/models"
	"github.com/dos-2/oddshub/sports"

	"github.com/rivo/tview"
)

// MastersGolf creates a slide for golf odds.
func MastersGolf(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	headerString := "Commencement Date|Teams|Players|Bookmaker|Outrights||\n"

	models.LoadEvent(string(sports.Golf_masters_tournament_winner), games)

	return "Masters", GetHeader(models.Golf_masters_tournament_winner), CreateRoundRobinTable(string(models.Golf_masters_tournament_winner), headerString, games)
}
