package main

import (
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals int, misses int, assists int) Player {
	player := Player{name, goals, misses, assists, 0}
	(&player).calculateRating()
	return player
}

func (p *Player) calculateRating() {
	switch p.Misses {
	case 0:
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2)
	default:
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	}
}

func (p *Player) gm() float64 {
	switch p.Misses {
	case 0:
		return float64(p.Goals)
	default:
		return float64(p.Goals) / float64(p.Misses)
	}

}

func nameCompare(p1 Player, p2 Player) int {
	n1 := []rune(p1.Name)
	n2 := []rune(p2.Name)
	for i := range min(len(n1), len(n2)) {
		if n1[i] > n2[i] {
			return 1
		}
		if n1[i] < n2[i] {
			return -1
		}
	}

	if len(n1) > len(n2) {
		return 1
	}
	if len(n1) < len(n2) {
		return -1
	}

	return 0
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(p1 Player, p2 Player) int {
		if p1.Goals > p2.Goals {
			return -1
		}
		if p1.Goals < p2.Goals {
			return 1
		}
		return nameCompare(p1, p2)
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(p1 Player, p2 Player) int {
		if p1.Rating > p2.Rating {
			return -1
		}
		if p1.Rating < p2.Rating {
			return 1
		}
		return nameCompare(p1, p2)
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(p1 Player, p2 Player) int {
		if (&p1).gm() > (&p2).gm() {
			return -1
		}
		if (&p1).gm() < (&p2).gm() {
			return 1
		}
		return nameCompare(p1, p2)
	})
	return players
}
