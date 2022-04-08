package mlog

import (
	"github.com/gdamore/tcell/v2"
	"golang.org/x/text/language"
)

type CodeBlock interface {
	// Identifier will be used in compiled code.
	Identifier() string
	Colour() int32
	ColourTview() int32

	DisplayName(lang language.Tag) string
	Description(lang language.Tag) string
}

var (
	pinkTview   = tcell.GetColor("pink").Hex()
	redTview    = tcell.GetColor("red").Hex()
	purpleTview = tcell.GetColor("purple").Hex()
	blueTview   = tcell.GetColor("blue").Hex()
	yellowTview = tcell.GetColor("yellow").Hex()
)

type CodeBlockPink struct {
}

func (CodeBlockPink) Colour() int32 {
	return 0xA08A8A
}

func (CodeBlockPink) ColourHex() int32 {
	return pinkTview
}

type CodeBlockRed struct {
}

func (CodeBlockRed) Colour() int32 {
	return 0xD4816B
}

func (CodeBlockRed) ColourHex() int32 {
	return redTview
}

type CodeBlockPurple struct {
}

func (CodeBlockPurple) Colour() int32 {
	return 0x877BAD
}

func (CodeBlockPurple) ColourHex() int32 {
	return purpleTview
}

type CodeBlockBlue struct {
}

func (CodeBlockBlue) Colour() int32 {
	return 0x6BB2B2
}

func (CodeBlockBlue) ColourHex() int32 {
	return blueTview
}

type CodeBlockYellow struct {
}

func (CodeBlockYellow) Colour() int32 {
	return 0xC7B59D
}

func (CodeBlockYellow) ColourHex() int32 {
	return yellowTview
}
