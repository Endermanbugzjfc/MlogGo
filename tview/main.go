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

	app := tview.NewApplication()
	root := tview.NewFlex().SetFullScreen(true)
	add := tview.NewFlex()
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
		AddItem("[pink]Read", "Read a number from a linked memory cell.", 'r', nil).
		AddItem("[pink]Draw", "", 'w', nil).
		AddItem("[red]Draw Flush", "Flush queued [yellow]Draw [green]operations to a displau.", 'f', nil).
		AddItem("[red]Get Link", "", 'g', nil).
		AddItem("[red]Radar", "", 'a', nil).
		AddItem("[purple]Set", "", 's', nil).
		AddItem("[blue]End", "", 'e', nil).
		AddItem("[yellow]Unit Bind", "", 'b', nil).
		AddItem("[yellow]Unit Radar", "", 'd', nil)

	addB.
		AddItem("[pink]Write", "Write a number to a linked memory cell.", 'w', nil).
		AddItem("[pink]Print", "", 't', nil).
		AddItem("[red]Print Flush", "", 'f', nil).
		AddItem("[red]Control", "", 'c', nil).
		AddItem("[red]Sensor", "", 's', nil).
		AddItem("[purple]Operation", "", 'a', nil).
		AddItem("[blue]Jump", "", 'g', nil).
		AddItem("[yellow]Unit Control", "", 'r', nil).
		AddItem("[yellow]Unit Locate", "", 'e', nil).
		Blur()

	if err := app.
		SetRoot(root, true).
		SetFocus(addA).
		Run(); err != nil {
		panic(err)
	}
}
