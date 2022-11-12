package main

import (
	"fmt"
	_ "github.com/f-sev/cryptowatcher/internal/currencies"
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

// инфо о курсе
// автозапуск
// хранить в env
// кнопка для автообновления
func onReady() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	Cron = cron.New()
	runCollectData()
	_, err := Cron.AddFunc("@every 5m", runCollectData)
	if err != nil {
		fmt.Printf("error while creating cron %s", err.Error())
	}

	Cron.Start()
	wg.Wait()
}

func runCollectData() {
	data.Manager.Collect()
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
