package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/endermanbugzjfc/mloggo/pkg/editor"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	version   = "1.0.0"
	corporate = `read result cell1 0
write result cell1 0
draw clear 0 1 2 0 0 0
draw color 0 1 2 255 0 0
draw stroke 0 1 2 255 0 0
draw line 0 1 2 255 0 0
draw rect 0 1 2 255 0 0
draw lineRect 0 1 2 255 0 0
draw poly 0 1 2 255 3 0
draw linePoly 0 1 2 255 3 0
draw triangle 0 1 2 255 3 4
draw image 0 1 @copper 32 2 0
print "frog"
drawflush display1
printflush message1
getlink result 0
control enabled block1 0 0 0 0
control shoot block1 0 1 2 0
control shootp block1 0 1 2 0
control configure block1 0 1 2 0
control color block1 0 1 2 0
radar enemy any player distance turret1 1 result
sensor result block1 @flarogus-lean
set result 0
op noise result a b
end
jump 25 equal x false
jump 25 notEqual x false
jump 25 lessThan x false
jump 25 lessThanEq x false
jump 25 greaterThan x false
jump 25 greaterThanEq x false
jump 25 strictEqual x false
jump 25 always x false
ubind @poly
ucontrol idle 0 0 0 0 0
ucontrol stop 0 0 0 0 0
ucontrol move 0 1 0 0 0
ucontrol approach 0 1 2 0 0
ucontrol boost 0 1 2 0 0
ucontrol pathfind 0 1 2 0 0
ucontrol target 0 1 2 0 0
ucontrol targetp 0 1 2 0 0
ucontrol itemDrop 0 1 2 0 0
ucontrol itemTake 0 1 2 0 0
ucontrol payDrop 0 1 2 0 0
ucontrol payTake 0 1 2 0 0
ucontrol mine 0 1 2 0 0
ucontrol flag 0 1 2 0 0
ucontrol build 0 1 2 3 4
ucontrol getBlock 0 1 2 3 4
ucontrol within 0 1 2 3 4
uradar enemy flying boss armor 0 1 result
ulocate ore core true @copper outx outy found building`
)

func makeHeader(text string) string {
	return fmt.Sprintf("【MLOG %s】%s", version, text)
}

func main() {
	editor.Init()
	logger := editor.GetLogger()
	logger.SetDebugMode(true)
	// configArgument := editor.RegisterConfigArgument()

	// flag.Parse()

	// logrusLogger := logrus.StandardLogger()
	// editor.InitLogrus(logrusLogger)
	// log := editor.LogrusToEditorLogger(logrusLogger)

	// editor.MustLoadConfig(log, *configArgument)
	// defaultKeys := editor.DefaultConfig().Keys
	// header := makeHeader("Make in Hong Kong \u1F1F")

	// box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	editor.NextKey(event.Rune(), config)

	// 	return event
	// })

	app := tview.NewApplication()
	defer func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	horizontal := tview.NewFlex().SetFullScreen(true)
	defer app.SetRoot(horizontal, true).SetFocus(horizontal)
	horizontal.SetBorder(true).SetTitleAlign(tview.AlignLeft)
	horizontalTitleStop := make(chan func(running bool) (stop bool))
	go marqueeTitle(app, horizontal.Box, horizontalTitleStop, true, "Read: ", `朋友是一個堅忍不拔的紀錄片 Lorem ipsum dolor sit amet`)

	editorView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	defer horizontal.AddItem(editorView, 0, 10, true)

	numSelections := 0
	go func() {
		for _, word := range strings.Split(corporate, " ") {
			if word == "the" {
				word = "[red]the[white]"
			}
			if word == "to" {
				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
				numSelections++
			}
			fmt.Fprintf(editorView, "%s ", word)
			// time.Sleep(200 * time.Millisecond)
		}
	}()
	editorView.SetDoneFunc(func(key tcell.Key) {
		currentSelection := editorView.GetHighlights()
		if key == tcell.KeyEnter {
			if len(currentSelection) > 0 {
				editorView.Highlight()
			} else {
				editorView.Highlight("0").ScrollToHighlight()
			}
		} else if len(currentSelection) > 0 {
			index, _ := strconv.Atoi(currentSelection[0])
			if key == tcell.KeyTab {
				index = (index + 1) % numSelections
			} else if key == tcell.KeyBacktab {
				index = (index - 1 + numSelections) % numSelections
			} else {
				return
			}
			editorView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
		}
	})
}
