/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package slides

import (
	"fmt"
	"oddshub/colors"
	"oddshub/models"
	"strings"
	// "time"
)

// FormatTournamentEvent formats a single event into a table-ready string.
func FormatTournamentEvent(event models.Event) string {
	var builder strings.Builder

	tournamentOdds := ExtractTournamentOdds(event)

	//	// Load time location once
	//	loc, err := time.LoadLocation("America/New_York")
	//	if err != nil {
	//		fmt.Println("Error loading time location:", err)
	//		return ""
	//	}

	// Format commencement time once
	commenceDate := event.CommenceTime.Format("01/02/2006")
	//commenceTime := event.CommenceTime.In(loc).Format("03:04 PM") + " ET"

	// Format bookmaker once
	bookmaker := fmt.Sprintf("[#FFFFFF:#333333]%s", event.Bookmakers[0].Title)

	// Iterate over player odds
	for _, playerOdd := range tournamentOdds {
		playerColors := getPlayerColors(event.SportKey, event.HomeTeam)                                     // Adjusted for F1 or future use
		teamText := fmt.Sprintf("[%s:%s]%s", playerColors.SecondaryColor, playerColors.PrimaryColor, "N/A") // Adjusted for F1 or future use
		playerText := fmt.Sprintf("[%s:%s]%s", playerColors.SecondaryColor, playerColors.PrimaryColor, playerOdd.Name)
		outrightOdd := formatMoneylineWithColor(playerOdd.Outright.Price)

		// Format player row
		playerRow := fmt.Sprintf("[%s:%s]%s|[%s:%s]%s|%s|%s|%s|%s|%s\n",
			playerColors.SecondaryColor, playerColors.PrimaryColor, commenceDate, playerColors.SecondaryColor, playerColors.PrimaryColor,
			teamText, playerText, bookmaker, outrightOdd, "[#FFFFFF:#333333]N/A", "[#FFFFFF:#333333]N/A")
		builder.WriteString(playerRow)
	}

	return builder.String()
}

// ExtractTournamentOdds extracts player odds from an event.
func ExtractTournamentOdds(event models.Event) []models.PlayerOdds {
	var tournamentOdds []models.PlayerOdds
	if len(event.Bookmakers) == 0 || len(event.Bookmakers[0].Markets) == 0 {
		return tournamentOdds
	}

	// Extracting from the first bookmaker and market
	market := event.Bookmakers[0].Markets[0]
	for _, outcome := range market.Outcomes {
		tournamentOdds = append(tournamentOdds, models.PlayerOdds{
			Name: outcome.Name,
			Outright: models.Outright{
				Price: outcome.Price,
			}})
	}
	return tournamentOdds
}

// getPlayerColors retrieves the team colors for a player or team.
func getPlayerColors(sport string, team string) models.TeamColors {
	colorsMap := colors.ColorsMap[sport]
	playerColors, exists := colorsMap[team]
	if !exists {
		playerColors = models.TeamColors{
			PrimaryColor:   "#333333",
			SecondaryColor: "#FFFFFF",
		}
	}
	return playerColors
}
