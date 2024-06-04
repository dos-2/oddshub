package models

// TeamOdds holds the odds information for both home and away teams.
type TeamOdds struct {
	HomeOdds Basic_Odds
	AwayOdds Basic_Odds
}

// Odds holds the spread, moneyline, and totals information.
type Basic_Odds struct {
	Spread    Spread
	Moneyline Moneyline
	Totals    Totals
}

// Spread holds spread betting information.
type Spread struct {
	Price float64
	Point float64
}

// Moneyline holds moneyline betting information.
type Moneyline struct {
	Price float64
}

// Totals holds totals betting information.
type Totals struct {
	OverPrice  float64
	OverPoint  float64
	UnderPrice float64
	UnderPoint float64
}
