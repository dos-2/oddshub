package ui

import (
  "oddshub/models"
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

func Formatter(eventsMap map[string][]models.Event) map[string]string {
  var pagesMap = make(map[string]string)
//
//  for sportName, events := range eventsMap {
//    pagesMap[sportName] = ""
//    for _, event := range events {
//      var row string = ""
//      row = append(row, event.CommenceTime)
//      row = append(row, "|HOME|")
//      row = append(row, event.HomeTeam + "|")
//      row  
//
//    }
//  }
    pagesMap["Americanfootball_ncaaf"] = tableData
    return pagesMap
}
