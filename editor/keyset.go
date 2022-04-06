package editor

import (
	"strings"

	"github.com/df-mc/atomic"
	"github.com/gdamore/tcell/v2"
)

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

func (keyset KeySet) MatchTcellEventKey(event *tcell.EventKey) bool {
	name := event.Name()
	split := strings.Split(name, "Rune[")
	name = split[1]

	split = strings.Split(name, "]")
	name = split[0]

	keysetString := string(keyset)
	if len(split) > 1 {
		return name == keysetString
	}

	name = strings.ToLower(name)
	keysetString = strings.ToLower(keysetString)
	return name == keysetString
}
