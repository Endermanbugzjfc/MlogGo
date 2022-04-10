package main

import (
	"time"

	"github.com/endermanbugzjfc/mloggo/pkg/editor"
	"github.com/rivo/tview"
)

// marqueeTitle should be called in its OWN GOROUTINE.
// The goroutine terminates when sync channel closes,
// sync function is nil or returns false.
// Sync channel will try to receive value
// at the END of each cycle.
func marqueeTitle(
	app *tview.Application,
	box *tview.Box,
	syncChannel <-chan func() bool,
	useInnerWidth bool,
	prefix, text string,
) {
	firstUpdate := make(chan struct{})
	app.QueueUpdate(func() {
		close(firstUpdate)
	})
	<-firstUpdate

	title := prefix + text
	titleWidth := tview.TaggedStringWidth(title)
	mbText := []rune(text)
	mbTextLength := len(mbText)
	logger := editor.GetLogger()

	const mbTextMin = 2
	if mbTextLength < mbTextMin {
		logger.Debugf(
			"Text is too short for marquee (%d runes, require at least %d): %s",
			mbTextLength,
			mbTextMin,
			text,
		)
		return
	}

	t := time.NewTicker(time.Second / 3) // TODO: User-changable.
	mbText = append(mbText, ' ')
	rolled := true
	for {
		var boxWidth int
		if useInnerWidth {
			_, _, boxWidth, _ = box.GetInnerRect()
		} else {
			_, _, boxWidth, _ = box.GetRect()
		}
		if titleWidth <= boxWidth {
			if rolled {
				rolled = false
				box.SetTitle(title)
				app.Draw()
			}
			continue
		}
		rolled = true

		mbText = append(mbText[1:], mbText[0])
		var (
			titleTrim      string
			titleTrimWidth int
			ok             bool
		)
		for trim := 0; trim < mbTextLength; trim++ {
			mbTextTrim := mbText[:mbTextLength-trim]
			titleTrim = prefix + string(mbTextTrim)
			titleTrimWidth = tview.TaggedStringWidth(titleTrim)
			if titleTrimWidth <= boxWidth {
				ok = true
				app.Draw()
				break
			}
		}

		if !ok {
			logger.Warnf(
				"Not enough space to fit the shortest trimmed marquee title: %s",
				titleTrim,
			)
			logger.Debugf(
				"Total width %d is greater than box wdith %d.",
				titleTrimWidth,
				boxWidth,
			)
		}
		box.SetTitle(titleTrim)

		select {
		case <-t.C:
		case syncFunc := <-syncChannel:
			if syncFunc == nil || !syncFunc() {
				t.Stop()

				return
			}
		}
	}
}
