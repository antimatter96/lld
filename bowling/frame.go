package main

import (
	"fmt"
	"strings"
)

// frame is used internally
// type frame struct {
// }

// type frame interface {
// 	score() int
// 	isStrike() bool
// 	isSpare() bool
// }

type Frame struct {
	rolls []string
}

func (f *Frame) CanBeNonLastFrame() bool {
	if f.rolls[2] != "" {
		return false
	}

	return true
}

func (f *Frame) CanBeLastFrame() bool {
	return true
}

func IsValidInput(s string) *Frame {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return nil
	}

	s = strings.ToLower(s)

	f := &Frame{}

	f.rolls = strings.Split(s, " ")

	//fmt.Println(ss, len(ss))
	fmt.Println(s, f.rolls)

	return nil
}

func IsValidFrame(frame int, input string) bool {

	return false
}
