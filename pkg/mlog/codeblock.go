package mlog

import (
	"github.com/df-mc/atomic"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/text/language"
)

var CodeBlocks atomic.Value[[]CodeBlock]

type CodeBlock interface {
	// Identifier will be used in compiled code.
	Identifier() string
	Colour() int32
	ColourForCommandLine() int32

	DisplayName(lang language.Tag) string
	Description(lang language.Tag) string
	Parse(parts []string) []CodeBlockPart
}

type CodeBlockPart interface {
	Label(lang language.Tag) string
	Note(lang language.Tag) string
}

var (
	pink   = tcell.GetColor("pink").Hex()
	red    = tcell.GetColor("red").Hex()
	purple = tcell.GetColor("purple").Hex()
	blue   = tcell.GetColor("blue").Hex()
	yellow = tcell.GetColor("yellow").Hex()
)

type CodeBlockPink struct {
}

func (CodeBlockPink) Colour() int32 {
	return 0xA08A8A
}

func (CodeBlockPink) ColourForCommandLine() int32 {
	return pink
}

type CodeBlockRed struct {
}

func (CodeBlockRed) Colour() int32 {
	return 0xD4816B
}

func (CodeBlockRed) ColourForCommandLine() int32 {
	return red
}

type CodeBlockPurple struct {
}

func (CodeBlockPurple) Colour() int32 {
	return 0x877BAD
}

func (CodeBlockPurple) ColourForCommandLine() int32 {
	return purple
}

type CodeBlockBlue struct {
}

func (CodeBlockBlue) Colour() int32 {
	return 0x6BB2B2
}

func (CodeBlockBlue) ColourForCommandLine() int32 {
	return blue
}

type CodeBlockYellow struct {
}

func (CodeBlockYellow) Colour() int32 {
	return 0xC7B59D
}

func (CodeBlockYellow) ColourForCommandLine() int32 {
	return yellow
}
