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
	// Set up event listeners for mouse events
	table := tview.NewTable().
		SetFixed(1, 1).
		SetBorders(false).
		SetSelectable(true, false)

	// Set up the table cells
	for row, line := range strings.Split(tableData, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := tcell.ColorWhite
			rowBgColor := tcell.ColorBlack
			align := tview.AlignCenter
			selectable := true

			if row == 0 {
				rowBgColor = tcell.NewRGBColor(57, 255, 20)
				color = tcell.ColorBlack
				align = tview.AlignCenter
				selectable = false
			}

			if column == 0 && row == 0 {
				rowBgColor = tcell.NewRGBColor(0, 255, 255)
			}

			tableCell := tview.NewTableCell(cell).
				SetTextColor(color).
				SetAlign(align).
				SetSelectable(selectable).
				SetExpansion(1).
				SetBackgroundColor(rowBgColor).
				SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).
					Background(tcell.NewRGBColor(0, 255, 255)))

			table.SetCell(row, column, tableCell)
		}
	}

	table.SetBorder(true).SetTitle(sportName)
	return table
}
