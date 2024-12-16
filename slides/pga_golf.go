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

func PGAGolf(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	headerString := "Commencement Date|Teams|Players|Bookmaker|Outrights||\n"

	models.LoadEvent(string(sports.Golf_pga_championship_winner), games)

	return "PGA", GetHeader(models.Golf_pga_tournament_winner), CreateRoundRobinTable(string(models.Golf_pga_tournament_winner), headerString, games)
}
