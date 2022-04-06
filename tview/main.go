package main

import (
	"fmt"

	"github.com/df-mc/atomic"
	"github.com/endermanbugzjfc/mloggo/editor"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

const (
	version = "1.0.0"
)

var (
	configAtomic atomic.Value[editor.Config]

	app       *tview.Application
	root, add *tview.Flex
)

func makeHeader(text string) string {
	return fmt.Sprintf("【MLOG %s】%s", version, text)
}

func main() {
	logrusLogger := logrus.StandardLogger()
	editor.InitLogrus(logrusLogger)
	log := editor.LogrusToEditorLogger(logrusLogger)
	config := editor.MustLoadConfig(log)
	configAtomic.Store(config)
	// header := makeHeader("Make in Hong Kong \u1F1F")

	// box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	editor.NextKey(event.Rune(), config)

	// 	return event
	// })

	app = tview.NewApplication()
	root = tview.
		NewFlex().
		SetFullScreen(true).
		SetDirection(tview.FlexRow)
	root.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	add = tview.NewFlex()
	root.AddItem(add, 0, 1, false)

	add.
		SetBorder(true).
		SetTitle("Add Action Block").
		SetTitleAlign(tview.AlignLeft)

	addA := tview.NewList()
	add.AddItem(addA, 0, 1, false)

	addB := tview.NewList()
	add.AddItem(addB, 0, 1, false)

	for _, box := range [2]*tview.List{
		addA,
		addB,
	} {
		func(box *tview.List) {
			colour := box.GetBackgroundColor()
			box.
				SetFocusFunc(func() {
					box.
						SetBackgroundColor(colour).
						SetTitle("(<Shift>W: Go to the other side.)").
						SetBorder(true)
				}).
				SetBlurFunc(func() {
					box.
						SetBackgroundColor(tcell.ColorDarkGray).
						SetTitle("")
				}).
				SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
					// TODO: Customise key in config.
					if event.Rune() == 'W' {
						if box == addA {
							app.SetFocus(addB)
						} else {
							app.SetFocus(addA)
						}
					}

					return event
				})
		}(box)
	}

	addA.
		AddItem("[pink]Read", "Read a number from a linked memory cell.", 'r', addCodeBlock(0)).
		AddItem("[pink]Draw", "", 'w', addCodeBlock(1)).
		AddItem("[red]Draw Flush", "Flush queued [yellow]Draw [green]operations to a displau.", 'f', addCodeBlock(2)).
		AddItem("[red]Get Link", "", 'g', addCodeBlock(3)).
		AddItem("[red]Radar", "", 'a', addCodeBlock(4)).
		AddItem("[purple]Set", "", 's', addCodeBlock(5)).
		AddItem("[blue]End", "", 'e', addCodeBlock(6)).
		AddItem("[yellow]Unit Bind", "", 'b', addCodeBlock(7)).
		AddItem("[yellow]Unit Radar", "", 'd', addCodeBlock(8))

	addB.
		AddItem("[pink]Write", "Write a number to a linked memory cell.", 'w', addCodeBlock(9)).
		AddItem("[pink]Print", "", 't', addCodeBlock(10)).
		AddItem("[red]Print Flush", "", 'f', addCodeBlock(11)).
		AddItem("[red]Control", "", 'c', addCodeBlock(12)).
		AddItem("[red]Sensor", "", 's', addCodeBlock(13)).
		AddItem("[purple]Operation", "", 'a', addCodeBlock(14)).
		AddItem("[blue]Jump", "", 'g', addCodeBlock(15)).
		AddItem("[yellow]Unit Control", "", 'r', addCodeBlock(16)).
		AddItem("[yellow]Unit Locate", "", 'e', addCodeBlock(17)).
		Blur()

	if err := app.
		SetRoot(root, true).
		SetFocus(addA).
		Run(); err != nil {
		panic(err)
	}
}

func addCodeBlock(codeBlockType int) func() {
	// TODO: Replace type with enum.
	asyncRemoveItemChannel := make(chan struct{})
	go func(asyncRemoveItemChannel chan struct{}) {
		<-asyncRemoveItemChannel
		root.RemoveItem(add)

		block := tview.NewFlex()
		block.SetBorder(true)
		switch codeBlockType {
		}

		block.SetTitleAlign(tview.AlignLeft)
		block.AddItem(tview.NewBox(), 0, 1, false)
		root.AddItem(block, 0, 1, false)
		app.SetFocus(block)
	}(asyncRemoveItemChannel)

	return func() {
		select {
		case asyncRemoveItemChannel <- struct{}{}:
		default:
		}
	}
}
