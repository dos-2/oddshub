package slides

import (
	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func NFL_football(nextSlide func()) (string, tview.Primitive) {
	tableData := `Commencement Date|Location|Teams|Bookmaker|Money Line|Spread|Total
1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
|AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108 
`
	return "NFL Football", CreateTable("NFL Football", tableData)
}
