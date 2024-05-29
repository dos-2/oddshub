package models

import "time"

// Event represents a sports event.
type Event struct {
  ID             string     `json:"id"`
  SportKey       string     `json:"sport_key"`
  SportTitle     string     `json:"sport_title"`
  CommenceTime   time.Time  `json:"commence_time"`
  HomeTeam       string     `json:"home_team"`
  AwayTeam       string     `json:"away_team"`
  Bookmakers     []Bookmaker `json:"bookmakers"`
}

// Define a struct for the bookmaker
type Bookmaker struct {
  Key         string    `json:"key"`
  Title       string    `json:"title"`
  LastUpdate  time.Time `json:"last_update"`
  Markets     []Market  `json:"markets"`
}

// Define a struct for the market
type Market struct {
  Key         string    `json:"key"`
  LastUpdate  time.Time `json:"last_update"`
  Outcomes    []Outcome `json:"outcomes"`
}

// Define a struct for the outcome
type Outcome struct {
  Name    string  `json:"name"`
  Price   float64 `json:"price"`
  Point   float64 `json:"point,omitempty"`
}
