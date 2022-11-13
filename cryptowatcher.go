package main

import (
	"fmt"
	"github.com/f-sev/cryptowatcher/internal/currencies"
	"github.com/f-sev/cryptowatcher/internal/data"
	"github.com/getlantern/systray"
	"github.com/robfig/cron/v3"
	"sync"
)

var Cron *cron.Cron
var IsExistQuit bool

func main() {
	systray.Run(onReady, onExit)
}

// автозапуск
// хранить в env
// кнопка для автообновления
func onReady() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	Cron = cron.New()
	runCollectData()
	_, err := Cron.AddFunc("@every 10m", runCollectData)
	if err != nil {
		fmt.Printf("error while creating cron %s", err.Error())
	}

	Cron.Start()
	wg.Wait()
}

func runCollectData() {
	go currencies.UpdateRates()
	data.Manager.Collect()
	addQuitButton()
}

func addQuitButton() {
	if !IsExistQuit {
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
		}()
		IsExistQuit = true
	}
}

func onExit() {
	defer Cron.Stop()
}
