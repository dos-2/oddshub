package ui 

import (
  "strings"
  "fmt"
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

const tableData = `Commencement Date|Location|Teams|Bookmaker|Money Line|Spread|Total
1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
|AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108 
 

  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108 
  

  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108 


  1/6/2017|HOME|Ohio State Buckeyes|Fanduel|-385|-9.5 -110|O 47.5 +112
  |AWAY|Michigan Wolverines|Fanduel|+300|+9.5 -110| U 47.5 -108 
  `

const tableBasic = `[green]func[white] [yellow]main[white]() {
  table := tview.[yellow]NewTable[white]().
  [yellow]SetFixed[white]([red]1[white], [red]1[white])
  [yellow]for[white] row := [red]0[white]; row < [red]40[white]; row++ {
  [yellow]for[white] column := [red]0[white]; column < [red]7[white]; column++ {
  color := tcell.ColorWhite
  [yellow]if[white] row == [red]0[white] {
  color = tcell.ColorYellow
  } [yellow]else[white] [yellow]if[white] column == [red]0[white] {
  color = tcell.ColorDarkCyan
  }
  align := tview.AlignLeft
  [yellow]if[white] row == [red]0[white] {
  align = tview.AlignCenter
  } [yellow]else[white] [yellow]if[white] column == [red]0[white] || column >= [red]4[white] {
  align = tview.AlignRight
  }
  table.[yellow]SetCell[white](row,
  column,
  &tview.TableCell{
  Text:  [red]"..."[white],
  Color: color,
  Align: align,
  })
  }
  }
  tview.[yellow]NewApplication[white]().
  [yellow]SetRoot[white](table, true).
  [yellow]Run[white]()
  }`

func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func Table() {
  table := tview.NewTable().
    SetFixed(1, 1).
    SetBorders(false).
    SetSelectable(true,true)

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
  table.SetBorder(true).SetTitle("Table")



  // Set up the application
  app := tview.NewApplication().
    SetRoot(table, true).
    SetFocus(table)

  // Enable mouse support (optional, if you want to support mouse interactions)
  table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Key() {
    case tcell.KeyEnter:
    // Handle Enter key if needed
  }
    return event
  })

  // Clear the screen and run the application
  fmt.Print("\033[H\033[2J")
  if err := app.Run(); err != nil {
    panic(err)
  }
}

// Table demonstrates the Table.
//func Table(nextSlide func()) (title string, content tview.Primitive) {
//	table := tview.NewTable().
//		SetFixed(1, 1)
//	for row, line := range strings.Split(tableData, "\n") {
//		for column, cell := range strings.Split(line, "|") {
//			color := tcell.ColorWhite
//			if row == 0 {
//				color = tcell.ColorYellow
//			} else if column == 0 {
//				color = tcell.ColorDarkCyan
//			}
//			align := tview.AlignLeft
//			if row == 0 {
//				align = tview.AlignCenter
//			} else if column == 0 || column >= 4 {
//				align = tview.AlignRight
//			}
//			tableCell := tview.NewTableCell(cell).
//				SetTextColor(color).
//				SetAlign(align).
//				SetSelectable(row != 0 && column != 0)
//			if column >= 1 && column <= 3 {
//				tableCell.SetExpansion(1)
//			}
//			table.SetCell(row, column, tableCell)
//		}
//	}
//	table.SetBorder(true).SetTitle("Table")
//
//	code := tview.NewTextView().
//		SetWrap(false).
//		SetDynamicColors(true)
//	code.SetBorderPadding(1, 1, 2, 0)
//
//	list := tview.NewList()
//
//	basic := func() {
//		table.SetBorders(false).
//			SetSelectable(false, false).
//			SetSeparator(' ')
//		code.Clear()
//		fmt.Fprint(code, tableBasic)
//	}
//
//	separator := func() {
//		table.SetBorders(false).
//			SetSelectable(false, false).
//			SetSeparator(tview.Borders.Vertical)
//		code.Clear()
//		fmt.Fprint(code, tableSeparator)
//	}
//
//	borders := func() {
//		table.SetBorders(true).
//			SetSelectable(false, false)
//		code.Clear()
//		fmt.Fprint(code, tableBorders)
//	}
//
//	selectRow := func() {
//		table.SetBorders(false).
//			SetSelectable(true, false).
//			SetSeparator(' ')
//		code.Clear()
//		fmt.Fprint(code, tableSelectRow)
//	}
//
//	selectColumn := func() {
//		table.SetBorders(false).
//			SetSelectable(false, true).
//			SetSeparator(' ')
//		code.Clear()
//		fmt.Fprint(code, tableSelectColumn)
//	}
//
//	selectCell := func() {
//		table.SetBorders(false).
//			SetSelectable(true, true).
//			SetSeparator(' ')
//		code.Clear()
//		fmt.Fprint(code, tableSelectCell)
//	}
//
//	list.ShowSecondaryText(false).
//		AddItem("Basic table", "", 'b', basic).
//		AddItem("Table with separator", "", 's', separator).
//		AddItem("Table with borders", "", 'o', borders).
//		AddItem("Selectable rows", "", 'r', selectRow).
//		AddItem("Selectable columns", "", 'c', selectColumn).
//		AddItem("Selectable cells", "", 'l', selectCell).
//		AddItem("Next slide", "", 'x', nextSlide)
//	list.SetBorderPadding(1, 1, 2, 2)
//
//	basic()
//
//	return "Table", tview.NewFlex().
//		AddItem(tview.NewFlex().
//			SetDirection(tview.FlexRow).
//			AddItem(list, 10, 1, true).
//			AddItem(table, 0, 1, false), 0, 1, true).
//		AddItem(code, 56, 1, false)
//}
//package ui 
//
//import (
//	"github.com/gdamore/tcell/v2"
//	"github.com/rivo/tview"
//)
//
//func UI() {
//box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
//	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
//		panic(err)
//	}
//  
//  table := tview.NewTable().SetFixed(1,1).SetSelectable(true, true)
//  for row := 0; row < 40; row++ {
//    for column := 0; column < 7; column++ {
//      color := tcell.ColorWhite
//      if row == 0 {
//        color = tcell.ColorYellow
//      } else if column == 0 {
//        color = tcell.ColorDarkCyan
//      }
//      align := tview.AlignLeft
//      if row == 0 {
//        align = tview.AlignCenter
//      } else if column == 0  || column >= 4{
//        align = tview.AlignRight
//      }
//      table.SetCell(row,
//        column,
//        &tview.TableCell{
//          Text: "...",
//          Color: color,
//          Align: align,
//          NotSelectable: row == 0 || column == 0,
//        })
//    }
//  }
//  tview.NewApplication().SetRoot(table, true).EnableMouse(true).Run()
//}
