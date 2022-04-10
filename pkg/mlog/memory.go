package mlog

import (
	"golang.org/x/text/language"
)

type Read struct {
	CodeBlockPink
}

func (Read) Identifier() string {
	return "read"
}

func (Read) DisplayName(lang language.Tag) string {
	panic("implement me") // TODO: Get display name from Mindustry language file.
}

func (Read) Description(lang language.Tag) string {
	panic("implement me") // TODO: Get description from Mindustry language file.
}

func (Read) Parse(parts []string) []CodeBlockPart {
	panic("implement me") // TODO: Parse()
}

type Write struct {
	CodeBlockPink
}

func (Write) Identifier() string {
	return "write"
}

func (Write) DisplayName(lang language.Tag) string {
	panic("implement me") // TODO: Get display name from Mindustry language file.
}

func (Write) Description(lang language.Tag) string {
	panic("implement me") // TODO: Get description from Mindustry language file.
}

func (Write) Parse(parts []string) []CodeBlockPart {
	panic("implement me") // TODO: Parse()
}
