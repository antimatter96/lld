package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

const (
	playerA int = iota // 0
	playerB int = iota // 1
)

type Game struct {
	N int

	shipIdCounter int
	shipIdMap     map[int]string

	BoardA [][]int // 3 states, 0 => Empty, -1 => Already attempted, +ve => Ship Present
	BoardB [][]int

	ShipsOfA map[int]bool
	ShipsOfB map[int]bool

	turnOf int
}

func main() {
	//g := InitGame(5)

	//_ = g.AddShip("SH-1", 2, 0, 4, 3, 3)
	///_ = g.AddShip("SH-2", 1, 0, 0, 5, 3)

	//g.Display()
	//g.Display2()

	//simulate(g)

	g2 := InitGame(11)
	_ = g2.AddShip("SH-1", 2, 0, 4, 8, 3)
	_ = g2.AddShip("SH-2", 1, 0, 0, 8, 2)

	g2.Display2()
	simulate2(g2)
}

func InitGame(n int) *Game {
	size := n + 1
	game := &Game{N: size, shipIdCounter: 1, turnOf: playerA}

	game.BoardA = make([][]int, size)
	game.BoardB = make([][]int, size)

	for i := 0; i < size; i++ {
		game.BoardA[i] = make([]int, size/2)
		game.BoardB[i] = make([]int, size/2)
	}

	game.shipIdMap = make(map[int]string)
	game.ShipsOfA = make(map[int]bool)
	game.ShipsOfB = make(map[int]bool)

	return game
}

func (game *Game) Display() {
	for i := 0; i < game.N; i++ {
		fmt.Println(game.BoardA[i], game.BoardB[i])
	}
}

func getShipCoordinates(size int, x, y int) [][]int {
	coordinates := make([][]int, 0, size*size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			coordinates = append(coordinates, []int{y - i, x + j})
		}
	}

	return coordinates
}

func (game *Game) AddShip(id string, size int, xA, yA int, xB, yB int) error {
	// Overlap logic missing

	yA = game.N - yA - 1
	yB = game.N - yB - 1
	xB = xB - game.N/2

	if xA < 0 || yA < 0 || xA+size > game.N/2 || yA+size > game.N {
		return errors.New("playerA ship out of bounds")
	}

	if xB < 0 || yB < 0 || xB+size > game.N || yB+size > game.N {
		return errors.New("playerB ship out of bounds")
	}

	game.shipIdMap[game.shipIdCounter] = id

	shipACoordinates := getShipCoordinates(size, xA, yA)
	for _, coordinate := range shipACoordinates {
		game.BoardA[coordinate[0]][coordinate[1]] = game.shipIdCounter
	}

	shipBCoordinates := getShipCoordinates(size, xB, yB)
	for _, coordinate := range shipBCoordinates {
		game.BoardB[coordinate[0]][coordinate[1]] = game.shipIdCounter
	}

	game.ShipsOfA[game.shipIdCounter] = true
	game.ShipsOfB[game.shipIdCounter] = true

	game.shipIdCounter++
	return nil
}

func (game *Game) MakeMove(player int, x, y int) (string, error) {
	if player != game.turnOf {
		return "", errors.New("Not your turn")
	}

	targetX := game.N - y - 1

	targetBoard := &(game.BoardB)
	targetShip := &(game.ShipsOfB)

	targetY := x - game.N/2

	if player == playerB {
		targetBoard = &(game.BoardA)
		targetShip = &(game.ShipsOfA)

		targetY = x
	}

	if targetY < 0 || targetY > game.N/2 || targetX < 0 || targetX > game.N {
		return "", errors.New("out of bounds")
	}

	//fmt.Println(targetBoard, targetShip)
	//fmt.Println(targetX, targetY)

	if (*targetBoard)[targetX][targetY] < 0 {
		return "", errors.New("Already tried")
	}

	result := "MISS"
	if (*targetBoard)[targetX][targetY] > 0 {
		shipId := (*targetBoard)[targetX][targetY]
		if (*targetShip)[shipId] {
			result = "HIT"
			delete(*targetShip, shipId)
		}
	}

	(*targetBoard)[targetX][targetY] = -1
	game.turnOf = (game.turnOf + 1) % 2

	return result, nil
}

func (game *Game) StartGame(p1 FireStrategy, p2 FireStrategy) {
	fmt.Println("Game Started")
	players := []int{playerA, playerB}
	stratergies := []FireStrategy{p1, p2}

	i := 0

	for {
		fmt.Println("---")
		game.PrintStatus()

		for {
			x, y, _ := stratergies[i](game)
			// Input error handling missing

			fmt.Println("Player", game.TurnOf()+1, "trying", x, y)
			result, err := game.MakeMove(players[i], x, y)
			if err == nil {
				fmt.Println(result)
				break
			} else {
				fmt.Println("Error:", err, "RETRYING")
			}
		}

		if game.Over() {
			fmt.Println("Game over", "Player", game.LastTurnOf()+1, "won")
			break
		}

		i = (i + 1) % len(players)
	}

	game.PrintStatus()
}

type FireStrategy func(*Game) (int, int, error)

func (game *Game) Over() bool {
	return len(game.ShipsOfA) == 0 || len(game.ShipsOfB) == 0
}

func (game *Game) LastTurnOf() int {
	return (game.turnOf + 1) % 2
}

func (game *Game) TurnOf() int {
	return game.turnOf
}

func (game *Game) PrintStatus() {
	fmt.Println("Player 1 Ships Remaining:", len(game.ShipsOfA))
	fmt.Println("Player 2 Ships Remaining:", len(game.ShipsOfB))

	if !game.Over() {
		fmt.Println("Player", game.turnOf+1, "turn")
	}
}

// RANDOM RUNNERS

func simulate(game *Game) {
	var err error
	var result string

	result, err = game.MakeMove(playerA, 3, 0)
	fmt.Println("Error:", err, "Result", result)
	game.PrintStatus()

	result, err = game.MakeMove(playerB, 1, 1)
	fmt.Println("Error:", err, "Result", result)
	game.PrintStatus()

	result, err = game.MakeMove(playerA, 4, 3)
	fmt.Println("Error:", err, "Result", result)
	game.PrintStatus()

	// result, err = game.MakeMove(playerB, 1, 1)
	// fmt.Println("Error:", err, "Result", result)
	// game.PrintStatus()
}

//
func getRandomGenerator(minX, minY, maxX, maxY int) FireStrategy {
	return func(game *Game) (int, int, error) {
		return rand.Intn(maxX-minX) + minX, rand.Intn(maxY-minY+1) + minY, nil
	}
}

//
func simulate2(game *Game) {
	n := game.N
	game.StartGame(getRandomGenerator(n/2, 0, n, 5), getRandomGenerator(0, 0, n/2, 5))
}

// Better display
func (game *Game) row(player, i int) string {
	str := &strings.Builder{}

	row := game.BoardA[i]
	if player == playerB {
		row = game.BoardB[i]
	}
	for i := 0; i < game.N/2; i++ {
		if row[i] == 0 {
			str.WriteString("....")
		} else if row[i] > 0 {
			str.WriteString(string([]byte(game.shipIdMap[row[i]])[0:4]))
		}
		str.WriteString(" ")
	}

	str.WriteString(" ")

	return str.String()
}

func (game *Game) Display2() {
	for i := 0; i < game.N; i++ {
		fmt.Println(game.row(0, i), " || ", game.row(1, i))
	}
}
