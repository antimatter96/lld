package baseGame

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	custom_errors "lld/play/custom_errors"
)

type game interface {
	player(int) string
	validateMove(string) error
	move(string, int) error
	Over() bool
}

type BaseStruct struct {
	game
	player1, player2 string
}

var (
	promptPlayerName                = "Player %d Name"
	moveRegex        *regexp.Regexp = regexp.MustCompile("[, .]+")
)

func (game *BaseStruct) InputNames() error {
	if game.player1 != "" && game.player2 != "" {
		return nil
	}

	var result string
	var err error

	if game.player1 == "" {
		result, err = (&promptui.Prompt{
			Label:    fmt.Sprintf(promptPlayerName, 1),
			Validate: validateNameNonEmpty,
		}).Run()
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return err
		}

		result = strings.TrimSpace(result)

		game.player1 = result
	}

	result, err = (&promptui.Prompt{
		Label:    fmt.Sprintf(promptPlayerName, 2),
		Validate: game.validatePlayer2Name,
	}).Run()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return err
	}

	game.player2 = result

	return nil
}

func (game *BaseStruct) validatePlayer2Name(input string) error {
	if err := validateNameNonEmpty(input); err != nil {
		return err
	}
	input = strings.TrimSpace(input)

	if game.player1 == input {
		return custom_errors.NameTaken
	}

	return nil
}

func MakeMove(game game, i int) error {
	var move string
	var err error

	for move, err = inputMove(game, i); err != nil; {
	}
	return game.move(move, i)
}

func inputMove(g game, i int) (string, error) {
	result, err := (&promptui.Prompt{
		Label:    g.player(i) + " Input move",
		Validate: g.validateMove,
	}).Run()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return "", err
	}

	return result, nil
}

func validateMoveNoContext(input string) error {
	split := moveRegex.Split(input, -1)

	if len(split) != 2 {
		return custom_errors.TooFew
	}

	if _, err := strconv.Atoi(split[0]); err != nil {
		return custom_errors.InvalidCharacters
	}
	if _, err := strconv.Atoi(split[1]); err != nil {
		return custom_errors.InvalidCharacters
	}

	return nil
}

func validateNameNonEmpty(input string) error {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return custom_errors.EmptyString
	}
	return nil
}
