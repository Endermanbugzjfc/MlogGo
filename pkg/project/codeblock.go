package project

import "golang.org/x/text/language"

type CodeBlockType interface {
	// Name is the display name.
	Name(language.Tag) string
	// Identifier is the name used in compiled code.
	Identifier() string
}
