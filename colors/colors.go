package colors

import (
	"oddshub/models"
	"oddshub/sports"
)

var ColorsMap = map[string]map[string]models.TeamColors{
	string(sports.Americanfootball_ncaaf):         NCAATeamColorMap,
	string(sports.Americanfootball_nfl):           NFLTeamColorMap,
	string(sports.Basketball_nba):                 NBATeamColorMap,
	string(sports.Golf_masters_tournament_winner): {},
	string(sports.Icehockey_nhl):                  NHLTeamColorMap,
	string(sports.Mma_mixed_martial_arts):         {},
	string(sports.Soccer_usa_mls):                 MLSTeamColorMap,
	string(sports.Tennis_atp_french_open):         {},
	string(sports.Tennis_wta_french_open):         {},
	string(sports.Baseball_mlb):                   MLBTeamColorMap,
	string(sports.Baseball_ncaa):                  NCAATeamColorMap,
	string(sports.Basketball_ncaa):                NCAATeamColorMap,
	string(sports.Boxing): {},
  string(sports.Soccer_brazil_campeonato):       CampeonatoTeamColorMap,
  string(sports.Soccer_epl):                     EuroTeamColorMap,
  string(sports.Cricket_ipl):                    IPLTeamColorMap,
  string(sports.Soccer_spain_la_liga):           EuroTeamColorMap,
  string(sports.Tennis_atp_wimbledon):           {},
  string(sports.Rugbyleague_nrl):                NRLTeamColorMap,
  string(sports.Golf_pga_championship_winner):   {},
  string(sports.Soccer_uefa_europa_league):      EuroTeamColorMap,
  string(sports.Tennis_wta_wimbledon):           {},
}
