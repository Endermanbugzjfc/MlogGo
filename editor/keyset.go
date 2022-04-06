package editor

import "github.com/df-mc/atomic"

type Mode int

const (
	ModeSuper Mode = iota
	ModeBlockSuper
	ModeBlockEdit
)

var (
	mode atomic.Value[Mode]
)

type KeySet string

func (keyset KeySet) MatchWithRune(toMatch rune) bool {
	return rune(keyset[0]) == toMatch
}

// TODO: Complex keysets support.
