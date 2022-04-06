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

//lint:ignore U1000 Function is used through compiler directives.
func Pink() int32 {
	return 0xA08A8A
}

//lint:ignore U1000 Function is used through compiler directives.
func PinkTview() int32 {
	return tcell.GetColor("pink").Hex()
}

//lint:ignore U1000 Function is used through compiler directives.
func Red() int32 {
	return 0xA08A8A
}

//lint:ignore U1000 Function is used through compiler directives.
func RedTview() int32 {
	return tcell.GetColor("pink").Hex()
}

//lint:ignore U1000 Function is used through compiler directives.
func Red() int32 {
	return 0xD4816B
}

//lint:ignore U1000 Function is used through compiler directives.
func RedTview() int32 {
	return tcell.GetColor("red").Hex()
}

//lint:ignore U1000 Function is used through compiler directives.
func Purple() int32 {
	return 0x877BAD
}

//lint:ignore U1000 Function is used through compiler directives.
func PurpleTview() int32 {
	return tcell.GetColor("purple").Hex()
}

//lint:ignore U1000 Function is used through compiler directives.
func Blue() int32 {
	return 0x6BB2B2
}

//lint:ignore U1000 Function is used through compiler directives.
func BlueTview() int32 {
	return tcell.GetColor("blue").Hex()
}

//lint:ignore U1000 Function is used through compiler directives.
func Yellow() int32 {
	return 0xC7B59D
}

//lint:ignore U1000 Function is used through compiler directives.
func YellowTview() int32 {
	return tcell.GetColor("yellow").Hex()
}
