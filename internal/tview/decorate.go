package main

import (
	"time"

	"github.com/endermanbugzjfc/mloggo/pkg/editor"
	"github.com/rivo/tview"
)

// marqueeTitle stops when sync channel closes,
// sync function is nil or returns false.
// Sync channel will try to receive value
// at the END of each cycle.
// Also should be called in its OWN GOROUTINE.
func marqueeTitle(
	app *tview.Application,
	sync <-chan func() bool,
	box *tview.Box,
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

	box.SetTitle(title)
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

	t := time.NewTicker(time.Second / 2) // TODO: User-changable.
	mbTextShift := mbText
	for {
		var boxWidth int
		if useInnerWidth {
			_, _, boxWidth, _ = box.GetInnerRect()
		} else {
			_, _, boxWidth, _ = box.GetRect()
		}
		if titleWidth <= boxWidth {
			continue
		}

		mbTextShift = append(mbTextShift[1:], mbTextShift[0])
		logger.Infof("%s", string(mbTextShift))
		var (
			titleNew string
			ok       bool
		)
		for trim := 0; trim < mbTextLength; trim++ {
			mbTextTrim := mbTextShift[:mbTextLength-trim]
			titleTrim := prefix + string(mbTextTrim)
			titleTrimLength := tview.TaggedStringWidth(titleTrim)
			if titleTrimLength <= boxWidth {
				titleNew = titleTrim
				ok = true
				app.Draw()
				break
			}
		}

		box.SetTitle(titleNew)
		if !ok {
			logger.Warnf(
				"Not enough space to fit marquee title: %s",
				titleNew,
			)
			logger.Debugf(
				"Total text width %d is greater than box wdith %d.",
				tview.TaggedStringWidth(titleNew),
				boxWidth,
			)
		}

		select {
		case <-t.C:
		case syncFunc := <-sync:
			if syncFunc == nil || !syncFunc() {
				t.Stop()

				return
			}
		}
	}
}