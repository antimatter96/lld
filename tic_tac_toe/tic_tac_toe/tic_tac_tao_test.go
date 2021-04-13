package tic_tac_toe

import "testing"

func TestValidateMoveNoContext(t *testing.T) {
	cases := []struct {
		input string
		err   error
	}{
		{"", errTooFew},
		{"", errTooFew},
		{"a a", errInvalidCharacters},
		{"1 a", errInvalidCharacters},
		{"a 1", errInvalidCharacters},
		{"-1 2", errOutOfRange},
		{"2 -1", errOutOfRange},
		{"-1 -1", errOutOfRange},
		{"1 1", nil},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			err := validateMoveNoContext(c.input)
			if err != c.err {
				t.Fatalf("Wanted %v, but got %v", c.err, err)
			}
		})
	}
}
