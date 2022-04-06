package mlog

import "golang.org/x/text/language"

type CodeBlock interface {
	// Identifier will be used in compiled code.
	Identifier() string
	DisplayName(lang language.Tag) string
	Colour() int32
	ColourTview() int32
}
