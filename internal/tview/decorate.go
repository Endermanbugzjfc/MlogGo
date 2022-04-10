package main

import (
	"time"

	"github.com/endermanbugzjfc/mloggo/pkg/editor"
	"github.com/rivo/tview"
)

func marqueeTitle(
	box *tview.Box,
	prefix, text string,
) (stop chan<- struct{}) {
	title := prefix + text
	titleWidth := tview.TaggedStringWidth(title)
	mbText := []rune(text)
	mbTextLength := len(mbText)
	notShort := mbTextLength > 1
	stopUnsafe := make(chan struct{})
	stop = stopUnsafe
	_, _, boxWidth, _ := box.GetRect()

	logger := editor.GetLogger()
	cannotFit := func(title string) {
		logger.Warnf(
			"Not enough space to fit marquee title: %s",
			title,
		)
		logger.Debugf(
			"Total text width %d is greater than box wdith: %d.",
			tview.TaggedStringWidth(title),
			boxWidth,
		)
	}

	if notShort && titleWidth > boxWidth {
		go func(stop chan struct{}) {
			t := time.NewTicker(time.Second) // TODO: User-changable.
			mbTextShift := mbText
			for {
				last := []rune{mbTextShift[mbTextLength-1]}
				allExceptLast := mbTextShift[:mbTextLength-1]
				mbTextShift = append(last, allExceptLast...)

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
						box.SetTitle(titleNew)
						break
					}
				}

				if !ok {
					cannotFit(titleNew)
				}

				select {
				case <-t.C:
				case <-stop:
					t.Stop()
					return
				}
			}
		}(stopUnsafe)
	} else if !notShort {
		cannotFit(title)
	}
	box.SetTitle(title)

	close(stop)
	return
}
