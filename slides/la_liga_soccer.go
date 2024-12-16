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

// LaLigaSoccer creates a slide for La Liga soccer odds.
func LaLigaSoccer(pages *tview.Pages, games []models.Event) (string, string, tview.Primitive) {
	headerString := "Commencement Date|Location|Teams|Bookmaker|Spread –|Money –|Total –\n"

	models.LoadEvent(string(sports.Soccer_spain_la_liga), games)

	return "La Liga", GetHeader(models.Soccer_spain_la_liga), CreateH2HTable(pages, string(models.Soccer_spain_la_liga), headerString, games)
}
