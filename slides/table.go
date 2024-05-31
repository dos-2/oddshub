package slides 

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//const tableData = `Commencement Date|Location|Teams|Bookmaker|Money Line|Spread|Total
//  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
//  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108
//
//
//  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
//  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108
//
//
//  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
//  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108
//
//
//  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
//  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108
//  `

func CreateTable(sportName string, tableData string) *tview.Table {
	table := tview.NewTable().
		SetFixed(1, 1).
		SetBorders(false).
		SetSelectable(true, true)

	// Set up the table cells
	for row, line := range strings.Split(tableData, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := tcell.ColorWhite
			if row == 0 {
				color = tcell.ColorYellow
			} else if column == 0 {
				color = tcell.ColorDarkCyan
			}
			align := tview.AlignCenter
			if row == 0 {
				align = tview.AlignCenter
			}
			tableCell := tview.NewTableCell(cell).
				SetTextColor(color).
				SetAlign(align).
				SetSelectable(true).
				SetExpansion(1)
			table.SetCell(row, column, tableCell)
		}
	}

	table.SetBorder(true).SetTitle(sportName)
	return table
}
