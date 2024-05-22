package main
import(
  "oddshub/API"
  "fmt"
)

func main() {
  printInfoMap(API.GetAllUpcomingEventsMap())
  //API.GetAllUpcomingEvents()
  //printInfo(API.GetAllUpcomingEvents())
}

func printInfo(events []API.Event) {

  for _, event := range events {

    fmt.Printf("ID: %s\n", event.ID)
    fmt.Printf("Sport: %s\n", event.SportTitle)
    fmt.Printf("Commence Time: %s\n", event.CommenceTime)
    fmt.Printf("Home Team: %s\n", event.HomeTeam)
    fmt.Printf("Away Team: %s\n", event.AwayTeam)
    fmt.Println("Bookmakers:")


    for _, bookmaker := range event.Bookmakers {
      fmt.Printf("  - Title: %s\n", bookmaker.Title)
      fmt.Println("    Markets:")
      for _, market := range bookmaker.Markets {
        fmt.Printf("      - Key: %s\n", market.Key)
        fmt.Printf("        Last Update: %s\n", market.LastUpdate)
        fmt.Println("        Outcomes:")
        for _, outcome := range market.Outcomes {
          fmt.Printf("          - Name: %s, Price: %f, Point: %f\n", outcome.Name, outcome.Price, outcome.Point)
        }
      }
    }

    fmt.Println()
  }
}


func printInfoMap(eventsMap map[string][]API.Event) {

  for sportName, events := range eventsMap {
    fmt.Println(sportName)
    for _, event := range events {

      fmt.Printf("ID: %s\n", event.ID)
      fmt.Printf("Sport: %s\n", event.SportTitle)
      fmt.Printf("Commence Time: %s\n", event.CommenceTime)
      fmt.Printf("Home Team: %s\n", event.HomeTeam)
      fmt.Printf("Away Team: %s\n", event.AwayTeam)
      fmt.Println("Bookmakers:")


      for _, bookmaker := range event.Bookmakers {
        fmt.Printf("  - Title: %s\n", bookmaker.Title)
        fmt.Println("    Markets:")
        for _, market := range bookmaker.Markets {
          fmt.Printf("      - Key: %s\n", market.Key)
          fmt.Printf("        Last Update: %s\n", market.LastUpdate)
          fmt.Println("        Outcomes:")
          for _, outcome := range market.Outcomes {
            fmt.Printf("          - Name: %s, Price: %f, Point: %f\n", outcome.Name, outcome.Price, outcome.Point)
          }
        }
      }

      fmt.Println()
    }
  }
}
