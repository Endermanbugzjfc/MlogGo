package main

import "strings"

// Credit: https://gist.github.com/sisteamnik/c866cb7eed264ea3408d
func mbSubstr(s string, from, length int) string {
	//create array like string view
	wb := []string{}
	wb = strings.Split(s, "")

	//miss nil pointer error
	to := from + length

	if to > len(wb) {
		to = len(wb)
	}

	if from > len(wb) {
		from = len(wb)
	}

	return strings.Join(wb[from:to], "")
}
