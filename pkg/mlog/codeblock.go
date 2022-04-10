package mlog

import (
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/text/language"
)

// RegisterCodeBlock returns the overrided code block
// if it has the same identifier as the given one.
// Default code blocks will ALWAYS be overriden
// by externally registered code blocks.
func RegisterCodeBlock(codeBlock CodeBlock) (override CodeBlock) {
	codeBlocksMu.Lock()
	defer codeBlocksMu.Unlock()

	identifier := codeBlock.Identifier()
	identifier = strings.ToLower(identifier)
	override = codeBlocks[identifier]
	if override != nil {
		for _, defaultCodeBlock := range defaultCodeBlocks {
			if override == defaultCodeBlock {
				return
			}
		}
	}
	codeBlocks[identifier] = codeBlock

	return
}

func FindCodeBlockByIdentifier(identifier string) (codeBlock CodeBlock) {
	// TODO: Match case argument?
	codeBlocksMu.RLock()
	defer codeBlocksMu.RUnlock()

	codeBlock = codeBlocks[identifier]

	return
}

// TODO: GetCodeBlocks()?

var (
	codeBlocks   = map[string]CodeBlock{}
	codeBlocksMu sync.RWMutex

	defaultCodeBlocks = [18]CodeBlock{
		Read{},
		Write{},
	}
)

type CodeBlock interface {
	// Identifier will be used in compiled code.
	Identifier() string
	Colour() int32
	ColourForCommandLine() int32

	DisplayName(lang language.Tag) string
	Description(lang language.Tag) string
	ParseParts(parts []string) []CodeBlockPart
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
