package models

// TeamOdds holds the odds information for both home and away teams.
type PlayerOdds struct {
  Name string
  Outright Outright
}

type Outright struct {
  Price float64
}
