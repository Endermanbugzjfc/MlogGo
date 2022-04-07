package main

import (
	"flag"
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
	configArgument := editor.GetConfigArgument()

	flag.Parse()

	logrusLogger := logrus.StandardLogger()
	editor.InitLogrus(logrusLogger)
	log := editor.LogrusToEditorLogger(logrusLogger)

	config := editor.MustLoadConfig(log, *configArgument)
	configAtomic.Store(config)
	defaultKeys := editor.DefaultConfig().Keys
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

	keys := config.Keys
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
						SetTitle(fmt.Sprintf("(%s: Go to the other side.)", keys.SwitchCodeBlockList)).
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
		AddItem("[pink]Read", "Read a number from a linked memory cell.", keys.Read.GetRune(log, rune(defaultKeys.Read[0])), addCodeBlock(0)).
		AddItem("[pink]Draw", "", keys.Draw.GetRune(log, rune(defaultKeys.Draw[0])), addCodeBlock(1)).
		AddItem("[red]Draw Flush", "Flush queued [yellow]Draw [green]operations to a displau.", keys.DrawFlush.GetRune(log, rune(defaultKeys.DrawFlush[0])), addCodeBlock(2)).
		AddItem("[red]Get Link", "", keys.GetLink.GetRune(log, rune(defaultKeys.GetLink[0])), addCodeBlock(3)).
		AddItem("[red]Radar", "", keys.Radar.GetRune(log, rune(defaultKeys.Radar[0])), addCodeBlock(4)).
		AddItem("[purple]Set", "", keys.Set.GetRune(log, rune(defaultKeys.Set[0])), addCodeBlock(5)).
		AddItem("[blue]End", "", keys.End.GetRune(log, rune(defaultKeys.End[0])), addCodeBlock(6)).
		AddItem("[yellow]Unit Bind", "", keys.UnitBind.GetRune(log, rune(defaultKeys.UnitBind[0])), addCodeBlock(7)).
		AddItem("[yellow]Unit Radar", "", keys.UnitRadar.GetRune(log, rune(defaultKeys.UnitRadar[0])), addCodeBlock(8))

	addB.
		AddItem("[pink]Write", "Write a number to a linked memory cell.", keys.Write.GetRune(log, rune(defaultKeys.Write[0])), addCodeBlock(9)).
		AddItem("[pink]Print", "", keys.Print.GetRune(log, rune(defaultKeys.Print[0])), addCodeBlock(10)).
		AddItem("[red]Print Flush", "", keys.PrintFlush.GetRune(log, rune(defaultKeys.PrintFlush[0])), addCodeBlock(11)).
		AddItem("[red]Control", "", keys.Control.GetRune(log, rune(defaultKeys.Control[0])), addCodeBlock(12)).
		AddItem("[red]Sensor", "", keys.Sensor.GetRune(log, rune(defaultKeys.Sensor[0])), addCodeBlock(13)).
		AddItem("[purple]Operation", "", keys.Operation.GetRune(log, rune(defaultKeys.Operation[0])), addCodeBlock(14)).
		AddItem("[blue]Jump", "", keys.Jump.GetRune(log, rune(defaultKeys.Jump[0])), addCodeBlock(15)).
		AddItem("[yellow]Unit Control", "", keys.UnitControl.GetRune(log, rune(defaultKeys.UnitControl[0])), addCodeBlock(16)).
		AddItem("[yellow]Unit Locate", "", keys.UnitLocate.GetRune(log, rune(defaultKeys.UnitLocate[0])), addCodeBlock(17)).
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
