package slides

import (
	"fmt"
	"oddshub/colors"
	"oddshub/models"
	"strings"
	"time"
)

// FormatTeamEvent formats a single event into a table-ready string.
func FormatTeamEvent(event models.Event) string {
	var builder strings.Builder
	teamOdds := ExtractTeamOdds(event)

	// Load time location
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading time location:", err)
		return ""
	}

	// Format commencement time once
	commenceDate := event.CommenceTime.Format("01/02/2006")
	commenceTime := event.CommenceTime.In(loc).Format("03:04 PM") + " ET"

	// Precompute colors
	homeColors, awayColors := getColors(event.SportKey, event.HomeTeam, event.AwayTeam)
	homeTeamText := fmt.Sprintf("[%s:%s]%s", homeColors.SecondaryColor, homeColors.PrimaryColor, event.HomeTeam)
	awayTeamText := fmt.Sprintf("[%s:%s]%s", awayColors.SecondaryColor, awayColors.PrimaryColor, event.AwayTeam)

	// Iterate over bookmakers to format data
	for _, bookmaker := range event.Bookmakers {
		var underPoints string = "N/A"
		var overPoints string = "N/A"
		if teamOdds.HomeOdds.Totals.OverPoint != 0 {
			overPoints = fmt.Sprintf("%g", teamOdds.HomeOdds.Totals.OverPoint)
		}
		if teamOdds.AwayOdds.Totals.UnderPoint != 0 {
			underPoints = fmt.Sprintf("%g", teamOdds.AwayOdds.Totals.UnderPoint)
		}

		var spreadExists bool = teamOdds.HomeOdds.Spread.Price != 0 && teamOdds.AwayOdds.Spread.Price != 0
		var moneylineExists bool = teamOdds.HomeOdds.Moneyline.Price != 0 && teamOdds.AwayOdds.Moneyline.Price != 0
		var totalsExists bool = teamOdds.HomeOdds.Totals.OverPrice != 0 && teamOdds.AwayOdds.Totals.UnderPrice != 0

		bookmakerName := fmt.Sprintf("[%s:%s]%s", "#FFFFFF", "#333333", bookmaker.Title)

		var overOdds string = "[#FFFFFF:#333333]N/A"
		var awayOdds string = "[#FFFFFF:#333333]N/A"
		var homeSpread string = "[#FFFFFF:#333333]N/A"
		var awaySpread string = "[#FFFFFF:#333333]N/A"
		var homeMoney string = "[#FFFFFF:#333333]N/A"
		var awayMoney string = "[#FFFFFF:#333333]N/A"

		if totalsExists {
			overOdds = fmt.Sprintf("[%s:%s]O %s %s", "#FFFFFF", "#333333", overPoints, 
				formatMoneylineWithColor(teamOdds.HomeOdds.Totals.OverPrice))
			awayOdds = fmt.Sprintf("[%s:%s]U %s %s", "#FFFFFF", "#333333", underPoints,
				formatMoneylineWithColor(teamOdds.AwayOdds.Totals.UnderPrice))
		}

		if spreadExists {
			homeSpread = formatWithSign(teamOdds.HomeOdds.Spread.Point) + " " + 
			formatMoneylineWithColor(teamOdds.HomeOdds.Spread.Price)
			awaySpread = formatWithSign(teamOdds.AwayOdds.Spread.Point) + " " +  
			formatMoneylineWithColor(teamOdds.AwayOdds.Spread.Price) 
		}

		if moneylineExists {
			homeMoney = formatMoneylineWithColor(teamOdds.HomeOdds.Moneyline.Price)
			awayMoney = formatMoneylineWithColor(teamOdds.AwayOdds.Moneyline.Price)
		}

		// Format home team data with spread first
		homeRow := fmt.Sprintf("[%s:%s]%s|[%s:%s]HOME|%s|%s|%s|%s|%s\n",
			homeColors.SecondaryColor, homeColors.PrimaryColor, commenceDate, homeColors.SecondaryColor, homeColors.PrimaryColor,
			homeTeamText, bookmakerName, homeSpread, homeMoney , overOdds)
		builder.WriteString(homeRow)

		// Format away team data with spread first
		awayRow := fmt.Sprintf("[%s:%s]%s|[%s:%s]AWAY|%s|%s|%s|%s|%s\n",
			awayColors.SecondaryColor, awayColors.PrimaryColor, commenceTime, awayColors.SecondaryColor, awayColors.PrimaryColor,
			awayTeamText, bookmakerName, awaySpread, awayMoney, awayOdds)
		builder.WriteString(awayRow)

		builder.WriteString("||||||\n")
	}

	return builder.String()
}

// formatWithSign formats a float64 with a plus sign if it's positive, or an empty string if it's zero.
func formatWithSign(value float64) string {
	return fmt.Sprintf("[#FFFFFF:#333333]"+"%+g", value)
}

// formatMoneylineWithColor formats the moneyline with color based on the sign of the odds.
func formatMoneylineWithColor(value float64) string {
	var color string
	if value > 0 {
		color = "#39FF14"
	} else if value < 0 {
		color = "#FF3A3A"
	} else {
		color = "#FFFFFF"
	}
	var bgColor string = "#333333"
	return fmt.Sprintf("[%s:%s]%+g", color, bgColor, value)
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
	colorsMap := colors.ColorsMap[sport]
	homeColors, homeExists := colorsMap[homeTeam]
	awayColors, awayExists := colorsMap[awayTeam]
	if !homeExists {
		homeColors.PrimaryColor = "#333333"
		homeColors.SecondaryColor = "#FFFFFF"
	}
	if !awayExists {
		awayColors.PrimaryColor = "#333333"
		awayColors.SecondaryColor = "#FFFFFF"
	}
	return homeColors, awayColors
}
