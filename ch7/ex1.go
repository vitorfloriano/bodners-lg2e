package main

import (
	
)

type Team struct {
	teamName   string
	playerName string
}

type League struct {
	Teams Team
	Wins  map[string]int
}

func (l League) MatchResult(firstTeam, secondTeam string, ftScore, stScore int) {
	
}

func (l League) Ranking() []Team.teamName {
	
}


