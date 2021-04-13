package tic_tac_toe

import (
	"fmt"
	"strings"
)

func ticTacToePrintUtility() func(i, j, val int) string {
	const colorOff string = "\033[0m" // Text Reset

	var paddedText = []string{
		"   ",
		" X ",
		" O ",
	}

	var colorCodes = []string{
		"0",
		"178",
		"208",
	}

	return func(i, j, val int) string {
		return fmt.Sprintf("\033[38;5;%sm%s%s", colorCodes[val], paddedText[val], colorOff)
	}
}

func (game *TicTacToe) PrintBoard() {
	cellPrinter := ticTacToePrintUtility()
	limit := len(game.board)
	nDashed := (2 * limit) + limit + 1 + (1 * limit)
	minStringSize := nDashed * ((2 * limit) - 1)

	var b strings.Builder
	b.Grow(minStringSize)

	b.WriteString(strings.Repeat("-", nDashed))
	b.WriteByte('\n')
	for i := 0; i < N; i++ {
		b.WriteString("|")
		for j := 0; j < N-1; j++ {
			b.WriteString(fmt.Sprintf("%s|", cellPrinter(i, j, game.board[i][j])))
		}
		b.WriteString(fmt.Sprintf("%s|\n", cellPrinter(i, i, game.board[i][N-1])))
		b.WriteString(strings.Repeat("-", nDashed))
		if i != N-1 {
			b.WriteByte('\n')
		}
	}

	fmt.Println(b.String())
}
