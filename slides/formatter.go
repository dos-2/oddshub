package slides

//
//import (
//	"fmt"
//	"oddshub/models"
//	"strings"
//)
//
//func FormatEvents(eventsMap map[string][]models.Event) map[string]string {
//	formattedData := make(map[string]string)
//
//	for sportName, events := range eventsMap {
//		formattedData[sportName] = formatEventsForSport(events)
//	}
//
//	return formattedData
//}
//
//func formatEventsForSport(events []models.Event) string {
//	var rows []string
//
//	// Header row
//	headerRow := "Commencement Date|Location|Teams|Bookmaker|Money Line|Spread|Total"
//	rows = append(rows, headerRow)
//
//	// Format each event
//	for _, event := range events {
//		row := formatEvent(event)
//		rows = append(rows, row)
//	}
//
//	// Join rows into a single string
//	return strings.Join(rows, "\n")
//}
//
//func formatEvent(event models.Event) string {
//	// Format commencement date
//	commenceDate := event.CommenceTime.Format("01/02/2006")
//
//	// Format home team data
//	homeRow := fmt.Sprintf("%s|HOME|%s|%s|%g|%s|%s",
//		commenceDate, event.HomeTeam, event.Bookmaker, event.MoneyLine, event.Spread, event.Total)
//	rows = append(rows, homeRow)
//
//	// Format away team data
//	awayRow := fmt.Sprintf("|AWAY|%s|%s|%g|%s|%s",
//		event.AwayTeam, event.Bookmaker, event.MoneyLine, event.Spread, event.Total)
//	rows = append(rows, awayRow)
//
//	// Join rows into a single string
//	return strings.Join(rows, "\n")
//}
