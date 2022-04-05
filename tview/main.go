package main

import (
	"fmt"

	"github.com/df-mc/atomic"
	"github.com/endermanbugzjfc/mloggo/editor"
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
	log := logrus.StandardLogger()
	editor.InitLogrus(log)
	config := editor.MustLoadConfig(log)
	configAtomic.Store(config)

	config.LaunchEditor()
	go func() {

	}()
	header := makeHeader("Make in Hong Kong \u1F1F")
	box := tview.NewBox().
		SetBorder(true).
		SetTitle(header)

	if err := tview.
		NewApplication().
		SetRoot(box, true).
		Run(); err != nil {
		panic(err)
	}
}
