package editor

type KeySet string

func (keyset KeySet) MatchWithRune(toMatch rune) bool {
	return rune(keyset[0]) == toMatch
}

// TODO: Complex keysets support.
