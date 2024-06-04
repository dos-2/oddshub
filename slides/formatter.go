package slides

import (
	"fmt"
	"oddshub/colors"
	"oddshub/models"
	"strings"
)

// FormatTeamEvent formats a single event into a table-ready string.
func FormatTeamEvent(event models.Event) string {
	var rows []string
	teamOdds := ExtractTeamOdds(event)
	// Format commencement time
	commenceTime := event.CommenceTime.Format("01/02/2006")
	// Iterate over bookmakers to format data
	for _, bookmaker := range event.Bookmakers {
		var underPoints string = ""
		var overPoints string = ""
		if teamOdds.HomeOdds.Totals.OverPoint != 0 {
			overPoints = fmt.Sprintf("%g", teamOdds.HomeOdds.Totals.OverPoint)
		}
		if teamOdds.AwayOdds.Totals.UnderPoint != 0 {
			overPoints = fmt.Sprintf("%g", teamOdds.AwayOdds.Totals.UnderPoint)
		}
		homeColors, awayColors := getColors(event.SportKey, event.HomeTeam, event.AwayTeam)
		var homeTeamText string = fmt.Sprintf("[%s:%s]%s", homeColors.SecondaryColor, homeColors.PrimaryColor, event.HomeTeam)
		var awayTeamText string = fmt.Sprintf("[%s:%s]%s", awayColors.SecondaryColor, awayColors.PrimaryColor, event.AwayTeam)
		bookmakerName := bookmaker.Key // Get the bookmaker name
		// Format home team data with spread first
		homeRow := fmt.Sprintf("%s|HOME|%s|%s|%s %s|%s|O %s %s",
			commenceTime, homeTeamText, bookmakerName,
			formatWithSign(teamOdds.HomeOdds.Spread.Point), formatWithSign(teamOdds.HomeOdds.Spread.Price),
			formatWithSign(teamOdds.HomeOdds.Moneyline.Price),
			overPoints, formatWithSign(teamOdds.HomeOdds.Totals.OverPrice))
		rows = append(rows, homeRow)
		// Format away team data with spread first
		awayRow := fmt.Sprintf("|AWAY|%s|%s|%s %s|%s|U %s %s",
			awayTeamText, bookmakerName,
			formatWithSign(teamOdds.AwayOdds.Spread.Point), formatWithSign(teamOdds.AwayOdds.Spread.Price),
			formatWithSign(teamOdds.AwayOdds.Moneyline.Price),
			underPoints, formatWithSign(teamOdds.AwayOdds.Totals.UnderPrice))
		rows = append(rows, awayRow)
		rows = append(rows, "\n")
	}

	// Join rows into a single string
	return strings.Join(rows, "\n")
}

// formatWithSign formats a float64 with a plus sign if it's positive, or an empty string if it's zero.
func formatWithSign(value float64) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf("%+g", value)
}

func ExtractTeamOdds(event models.Event) models.TeamOdds {
	var teamOdds models.TeamOdds

	for _, bookmaker := range event.Bookmakers {
		for _, market := range bookmaker.Markets {
			switch market.Key {
			case "spreads":
				for _, outcome := range market.Outcomes {
					if outcome.Name == event.HomeTeam {
						teamOdds.HomeOdds.Spread.Price = outcome.Price
						teamOdds.HomeOdds.Spread.Point = outcome.Point
					} else if outcome.Name == event.AwayTeam {
						teamOdds.AwayOdds.Spread.Price = outcome.Price
						teamOdds.AwayOdds.Spread.Point = outcome.Point
					}
				}
			case "h2h":
				for _, outcome := range market.Outcomes {
					if outcome.Name == event.HomeTeam {
						teamOdds.HomeOdds.Moneyline.Price = outcome.Price
					} else if outcome.Name == event.AwayTeam {
						teamOdds.AwayOdds.Moneyline.Price = outcome.Price
					}
				}
			case "totals":

				for _, outcome := range market.Outcomes {
					if outcome.Name == "Over" {
						teamOdds.HomeOdds.Totals.OverPrice = outcome.Price
						teamOdds.HomeOdds.Totals.OverPoint = outcome.Point
					} else if outcome.Name == "Under" {
						teamOdds.AwayOdds.Totals.UnderPrice = outcome.Price
						teamOdds.AwayOdds.Totals.UnderPoint = outcome.Point
					}
				}
			}
		}
	}

	return teamOdds
}

func getColors(sport string, homeTeam string, awayTeam string) (models.TeamColors, models.TeamColors) {
	colorsMap, exists := colors.ColorsMap[sport]
	if !exists {
		return models.TeamColors{}, models.TeamColors{}
	}
	homeColors, homeExists := colorsMap[homeTeam]
	awayColors, awayExists := colorsMap[awayTeam]
	if !homeExists {
		homeColors.PrimaryColor = "[black]"
		homeColors.SecondaryColor = "[white]"
	}
	if !awayExists {
		awayColors.PrimaryColor = "[black]"
		awayColors.SecondaryColor = "[white]"
	}
	return homeColors, awayColors
}
