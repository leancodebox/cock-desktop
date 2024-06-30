package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/leancodebox/cock-desktop/resource"
	"github.com/leancodebox/cock-desktop/tm"
	"github.com/leancodebox/cock/jobmanager"
	"github.com/leancodebox/cock/jobmanagerserver"
	"log"
	"log/slog"
	"os"
)

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		stop()
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func main() {
	a := app.New()
	logLifecycle(a)
	w := a.NewWindow("cock-desktop")
	a.SetIcon(resource.GetLogo())
	a.Settings().SetTheme(&tm.MyTheme{})
	hello := widget.NewLabel("欢迎使用cock-desktop!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("你好", func() {
			hello.SetText("Welcome :)")
		}),
	))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	// 桌面系统设置托盘
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("bigBrother",
			fyne.NewMenuItem("关于", func() {
				w.Show()
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
	w.Resize(fyne.NewSize(300, 80))
	w.SetFixedSize(true)
	startCockServer()
	w.ShowAndRun()
}

func startCockServer() {
	fileData, err := os.ReadFile("jobConfig.json")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	jobmanager.Reg(fileData)
	jobmanagerserver.ServeRun()

}

func stop() {
	jobmanagerserver.ServeStop()
}
