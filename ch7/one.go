package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	teams []Team
	wins  map[string]int
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	if firstTeamScore > secondTeamScore {
		l.wins[firstTeamName]++
	} else {
		l.wins[secondTeamName]++
	}
}

func (l *League) Ranking() []string {
	var teamNames []string
	for name := range l.wins {
		teamNames = append(teamNames, name)
	}

	sort.Slice(teamNames, func(i, j int) bool {
		return l.wins[teamNames[i]] > l.wins[teamNames[j]]
	})

	return teamNames
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	res := r.Ranking()
	for _, v := range res {
		io.WriteString(w, v)
		io.WriteString(w, " ")
	}
}

func one() {

	league := League{
		teams: []Team{
			{name: "Lions", players: []string{"A", "B"}},
			{name: "Tigers", players: []string{"C", "D"}},
			{name: "Bears", players: []string{"E", "F"}},
			{name: "Wolves", players: []string{"G", "H"}},
			{name: "Eagles", players: []string{"I", "J"}},
		},
		wins: make(map[string]int),
	}

	// initialize all teams with 0 wins so they show in ranking
	for _, t := range league.teams {
		league.wins[t.name] = 0
	}

	// Each team plays each other twice (no draws, everyone wins at least once)
	results := []struct {
		t1 string
		s1 int
		t2 string
		s2 int
	}{
		{"Lions", 85, "Tigers", 80},
		{"Lions", 70, "Bears", 90},
		{"Lions", 95, "Wolves", 88},
		{"Lions", 60, "Eagles", 75},

		{"Tigers", 100, "Bears", 95},
		{"Tigers", 88, "Wolves", 90},
		{"Tigers", 92, "Eagles", 85},
		{"Bears", 85, "Wolves", 80},
		{"Bears", 82, "Eagles", 78},
		{"Wolves", 89, "Eagles", 92},

		// reverse fixtures (second round)
		{"Tigers", 84, "Lions", 91},
		{"Bears", 97, "Lions", 94},
		{"Wolves", 85, "Lions", 83},
		{"Eagles", 96, "Lions", 92},

		{"Bears", 91, "Tigers", 88},
		{"Wolves", 93, "Tigers", 90},
		{"Eagles", 87, "Tigers", 86},
		{"Wolves", 80, "Bears", 83},
		{"Eagles", 88, "Bears", 85},
		{"Eagles", 91, "Wolves", 89},
	}

	for _, r := range results {
		league.MatchResult(r.t1, r.s1, r.t2, r.s2)
	}

	RankPrinter(&league, os.Stdout)
}
