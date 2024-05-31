package ui 

import (
  "strings"
  "fmt"
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
  "oddshub/models"
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

func UI(eventsMap map[string][]models.Event) {
  var footballPage = Formatter(eventsMap) 
  table := CreateTable("AmericanFootball_ncaaf", footballPage["Americanfootball_ncaaf"]) 
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
