package main

import (
	"errors"
)

type PlayerGame struct {
	name   string
	frames []*Frame
	scores []int
}

var scoreMap = map[string]int{
	"1": 1, "2": 2, "3": 3,
	"4": 4, "5": 5, "6": 6,
	"7": 7, "8": 8, "9": 9,
	"X": 10,
	"-": 0,
}

func (pg *PlayerGame) ScoreTill(i int) (int, error) {
	allRolls := make([]string, 21)

	for _, frame := range pg.frames {
		allRolls = append(allRolls, frame.rolls...)
	}

	score := 0
	balls := 2

	for balls >= 0 {
		balls--
		if allRolls[i] == "X" {
			score += 10
			i += 2
			balls += 2
			continue
		} else if allRolls[i] == "/" {
			score -= scoreMap[allRolls[i-1]]
			score += 10
			i += 2
			balls += 2
		} else {
			score += scoreMap[allRolls[i]]
			i++
		}
	}
	return score, errors.New("")
}

// CalculateScore
func CalculateScore(a []int) int {
	frame := 1
	score := 0
	i := 0

	for frame < 11 {
		if a[i] == 10 {
			score += (10 + a[i+1] + a[i+2])
			i += 1
		} else if a[i]+a[i+1] == 10 {
			score += (10 + a[i+2])
			i += 2
		} else {
			score += (a[i] + a[i+1])
			i += 2
		}
		frame += 1
	}

	return score
}
