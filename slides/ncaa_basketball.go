package slides

import (
	"oddshub/models"
  "fmt"
	"github.com/rivo/tview"
)

// NcaaBasketball creates a slide for basketball odds.
func NCAABasketball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game)
	}
  fmt.Print(tableData)
	return "NCAA Basketball", GetHeader(models.Basketball_ncaa), CreateH2HTable("NCAA Basketball", tableData)
}
