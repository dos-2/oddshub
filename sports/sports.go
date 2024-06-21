/*
 * Copyright (c) 2024 dos-2
 * All rights reserved.
 */
package sports

type Sport string

const (
	Americanfootball_ncaaf            Sport = "americanfootball_ncaaf"
	Americanfootball_nfl              Sport = "americanfootball_nfl"
	Baseball_mlb                      Sport = "baseball_mlb"
	Baseball_ncaa                     Sport = "baseball_ncaa"
	Basketball_nba                    Sport = "basketball_nba"
	Basketball_ncaa                   Sport = "basketball_ncaab"
	Boxing                            Sport = "boxing_boxing"
	Cricket_ipl                       Sport = "cricket_ipl"
	Golf_masters_tournament_winner    Sport = "golf_masters_tournament_winner"
	Golf_pga_championship_winner      Sport = "golf_pga_championship_winner"
	Icehockey_nhl                     Sport = "icehockey_nhl"
	Mma_mixed_martial_arts            Sport = "mma_mixed_martial_arts"
	Rugbyleague_nrl                   Sport = "rugbyleague_nrl"
	Soccer_brazil_campeonato          Sport = "soccer_brazil_campeonato"
	Soccer_epl                        Sport = "soccer_epl"
	Soccer_spain_la_liga              Sport = "soccer_spain_la_liga"
	Soccer_uefa_europa_league         Sport = "soccer_uefa_europa_league"
	Soccer_usa_mls                    Sport = "soccer_usa_mls"
	Tennis_atp_french_open            Sport = "tennis_atp_french_open"
	Tennis_atp_wimbledon              Sport = "tennis_atp_wimbledon"
	Tennis_wta_french_open            Sport = "tennis_wta_french_open"
	Tennis_wta_wimbledon              Sport = "tennis_wta_wimbledon"
	Soccer_uefa_european_championship Sport = "soccer_uefa_european_championship"
	Soccer_conmebol_copa_america      Sport = "soccer_conmebol_copa_america"
)

func AllSports() []Sport {
	return []Sport{
		Americanfootball_ncaaf,
		Americanfootball_nfl,
		Baseball_mlb,
		Baseball_ncaa,
		Basketball_nba,
		Basketball_ncaa,
		Boxing,
		Cricket_ipl,
		Golf_masters_tournament_winner,
		Golf_pga_championship_winner,
		Icehockey_nhl,
		Mma_mixed_martial_arts,
		Rugbyleague_nrl,
		Soccer_brazil_campeonato,
		Soccer_epl,
		Soccer_spain_la_liga,
		Soccer_uefa_europa_league,
		Soccer_usa_mls,
		Soccer_uefa_european_championship,
		Soccer_conmebol_copa_america,
	}
}

// GetSportsMap returns a map where the keys are the string representations
// of the sports and the values are the corresponding Sport constants.
func GetSportsMap() map[string]Sport {
	sportsMap := make(map[string]Sport)
	for _, sport := range AllSports() {
		sportsMap[string(sport)] = sport
	}
	return sportsMap
}
