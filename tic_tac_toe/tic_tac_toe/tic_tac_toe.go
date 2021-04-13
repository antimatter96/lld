package tic_tac_toe

import (
	"errors"
	"fmt"
	base "lld/play/base_game"
	"lld/play/custom_errors"
	"regexp"
	"strconv"
)

const N = 3

var (
	errOutOfRange error          = errors.New("out of range [1-3]")
	moveRegex     *regexp.Regexp = regexp.MustCompile("[, .]+")
)

type TicTacToe struct {
	base.BaseStruct
	board [][]int
}

func NewTicTacToe() *TicTacToe {
	ttt := &TicTacToe{}
	ttt.board = make([][]int, N)
	for i := 0; i < N; i++ {
		ttt.board[i] = make([]int, N)
	}

	return ttt
}

func (game *TicTacToe) MakeBoard(player string) error {
	fmt.Println("created board for player", player)
	return nil
}

func (game *TicTacToe) validateMove(input string) error {
	if err := base.validateMoveNoContext(input); err != nil {
		return err
	}

	split := moveRegex.Split(input, -1)
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	if x < 0 || y < 0 {
		return errOutOfRange
	}

	if x > N || y > N {
		return errOutOfRange
	}

	if game.board[x][y] != 0 {
		return custom_errors.AlreadyPicked
	}

	return nil
}

func (game *TicTacToe) move(input string, player int) error {
	split := moveRegex.Split(input, -1)
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	game.board[x][y] = player

	return nil
}

func (game *TicTacToe) player(i int) string {
	if i == 1 {
		return game.player1 + `'s turn (X)`
	}
	return game.player2 + `'s turn (O)`
}

func (game *TicTacToe) over() bool {
	for i := 0; i < 3; i++ {
		if game.board[i][0] != 0 &&
			game.board[i][0] == game.board[i][1] &&
			game.board[i][0] == game.board[i][2] {
			return true
		}

		if game.board[0][i] != 0 &&
			game.board[0][i] == game.board[1][i] &&
			game.board[0][i] == game.board[2][i] {
			return true
		}
	}

	if game.board[0][0] != 0 &&
		game.board[0][0] == game.board[1][1] &&
		game.board[0][0] == game.board[2][2] {
		return true
	}
	if game.board[0][2] != 0 &&
		game.board[0][2] == game.board[1][1] &&
		game.board[0][2] == game.board[2][0] {
		return true
	}

	return false
}
