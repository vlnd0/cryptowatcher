package main

import (
	_ "github.com/f-sev/cryptowatcher/internal/currencies"
	"github.com/f-sev/cryptowatcher/internal/data"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	data.Manager.Collect()
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	// clean up here
}
