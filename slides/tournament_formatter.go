package slides

import (
	"fmt"
  "oddshub/models"
  "time"
  "oddshub/colors"
  "strings"
)

// FormatTeamEvent formats a single event into a table-ready string.
func FormatTournamentEvent(event models.Event) string {
	var rows []string
	tournamentOdds := ExtractTournamentOdds(event)
	// Format commencement time
  loc, err := time.LoadLocation("America/New_York")
  if err != nil {
    fmt.Print("Something wrong with the time.LoadLocation function")
  }
  commenceDate := event.CommenceTime.Format("01/02/2006")
  commenceTime := event.CommenceTime.In(loc).Format("03:04 PM")
  commenceTime += " ET"
  var bookmaker string = fmt.Sprintf("[#FFFFFF:#333333]%s",event.Bookmakers[0].Title)
  for _, playerOdd := range tournamentOdds {
    playerColors := getPlayerColors(event.SportKey, event.HomeTeam,) // later on for f1
    var teamText string = fmt.Sprintf("[%s:%s]%s", playerColors.SecondaryColor, playerColors.PrimaryColor, "N/A") // later for f1
    var playerText string = fmt.Sprintf("[%s:%s]%s", playerColors.SecondaryColor, playerColors.PrimaryColor, playerOdd.Name)
    var outrightOdd string = formatMoneylineWithColor(playerOdd.Outright.Price) 
    playerRow := fmt.Sprintf("[%s:%s]%s|[%s:%s]%s|%s|%s|%s|%s|%s",
      playerColors.SecondaryColor, playerColors.PrimaryColor, commenceDate, playerColors.SecondaryColor, playerColors.PrimaryColor,
      teamText, playerText, bookmaker, outrightOdd, "[#FFFFFF:#333333]N/A","[#FFFFFF:#333333]N/A")
    rows = append(rows, playerRow)
  }

  return strings.Join(rows, "\n")

}

func ExtractTournamentOdds(event models.Event) []models.PlayerOdds {
    var tournamentOdds []models.PlayerOdds
    if len(event.Bookmakers) == 0 {
        return tournamentOdds
    }
    var bookmaker models.Bookmaker = event.Bookmakers[0]
    if len(bookmaker.Markets) == 0 {
        return tournamentOdds
    }
    var market models.Market = bookmaker.Markets[0]
    for _, outcome := range market.Outcomes {
        var playerOdds models.PlayerOdds
        playerOdds.Name = outcome.Name
        playerOdds.Outright.Price = outcome.Price
        tournamentOdds = append(tournamentOdds, playerOdds)
    }
    return tournamentOdds
}

func getPlayerColors(sport string, team string) (models.TeamColors) {
  colorsMap := colors.ColorsMap[sport]
  playerColors, colorsExists := colorsMap[team]
  if !colorsExists{
    playerColors.PrimaryColor = "#333333"
    playerColors.SecondaryColor = "#FFFFFF"
  }
  return playerColors 
}
