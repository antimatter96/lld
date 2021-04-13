package tic_tac_toe

import (
	"fmt"
	base "lld/play/base_game"
)

func TicTacToeRunner() {
	game := NewTicTacToe()
	for err := game.InputNames(); err != nil; {
	}

	for i := 0; i < 2; i++ {
		for err := game.MakeBoard(fmt.Sprint(i + 1)); err != nil; {
		}
	}

	for over := game.Over(); !over; {
		for i := 0; i < 2; i++ {
			for err := base.MakeMove(game, i+1); err != nil; {
			}
			game.PrintBoard()
		}
	}

}
