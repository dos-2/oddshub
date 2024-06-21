/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package models

type Title string

const (
	Americanfootball_ncaaf         Title = "NCAA Football"
	Americanfootball_nfl           Title = "NFL Football"
	Baseball_mlb                   Title = "MLB Baseball"
	Baseball_ncaa                  Title = "NCAA Baseball"
	Basketball_ncaa                Title = "NCAA Basketball"
	Basketball_nba                 Title = "NBA Basketball"
	Boxing                         Title = "Boxing"
	Cricket_ipl                    Title = "IPL Cricket"
	Golf_masters_tournament_winner Title = "Masters Tournament"
	Golf_pga_tournament_winner     Title = "PGA Tour Champonship"
	Icehockey_nhl                  Title = "NHL Hockey"
	Mma_mixed_martial_arts         Title = "MMA"
	Rugby_irl                      Title = "NRL Rugby"
	Soccer_brazil_campeonato       Title = "Brazil Campeonato"
	Soccer_epl                     Title = "English Premier League"
	Soccer_spain_la_liga           Title = "Spain La Liga"
	Soccer_uefa_europa_league      Title = "UEFA Europa League"
	Soccer_usa_mls                 Title = "MLS Soccer"
	Tennis_atp_french_open         Title = "Mens French Open"
	Tennis_atp_wimbledon           Title = "Mens Wimbledon"
	Tennis_wta_french_open         Title = "Womens French Open"
	Tennis_wta_wimbledon           Title = "Womens Wimbledon"
	Soccer_uefa_european_champion  Title = "UEFA European Championship"
	Soccer_conmebol_copa_america   Title = "Copa America"
)

func AllTitles() []Title {
	return []Title{
		Americanfootball_ncaaf,
		Americanfootball_nfl,
		Baseball_mlb,
		Baseball_ncaa,
		Basketball_ncaa,
		Basketball_nba,
		Boxing,
		Golf_masters_tournament_winner,
		Icehockey_nhl,
		Mma_mixed_martial_arts,
		Soccer_brazil_campeonato,
		Soccer_epl,
		Soccer_spain_la_liga,
		Soccer_uefa_europa_league,
		Soccer_usa_mls,
		Tennis_atp_french_open,
		Tennis_wta_french_open,
		Tennis_atp_wimbledon,
		Tennis_wta_wimbledon,
	}
}
