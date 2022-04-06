package mlog

import (
	"github.com/gdamore/tcell/v2"
	"golang.org/x/text/language"
)

type CodeBlock interface {
	// Identifier will be used in compiled code.
	Identifier() string
	DisplayName(lang language.Tag) string
	Colour() int32
	ColourTview() int32
}

const (
	pink   = 0xA08A8A
	red    = 0xD4816B
	purple = 0x877BAD
	blue   = 0x6BB2B2
	yellow = 0xC7B59D
)

var (
	pinkTview   = tcell.GetColor("pink").Hex()
	redTview    = tcell.GetColor("red").Hex()
	purpleTview = tcell.GetColor("purple").Hex()
	blueTview   = tcell.GetColor("blue").Hex()
	yellowTview = tcell.GetColor("yellow").Hex()
)
