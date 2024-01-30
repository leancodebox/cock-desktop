package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/leancodebox/cock-desktop/tm"
	"github.com/spf13/cast"
	"log"
)

func logLifecycle(a fyne.App) {
	return
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func main() {
	a := app.NewWithID("github.com/leancodebox/cock-desktop")
	//makeTray(a)
	logLifecycle(a)
	w := a.NewWindow("cock-desktop")
	w.SetMaster()
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	// 桌面系统设置托盘
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("cock-desktop",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	//设置主体
	a.Settings().SetTheme(&tm.MyTheme{})

	header := []fyne.CanvasObject{
		widget.NewLabel("jobId"),
		widget.NewLabel("run"),
		widget.NewLabel("status"),
		widget.NewLabel("opt"),
	}

	var jobList []fyne.CanvasObject

	jobList = append(jobList, container.NewGridWithColumns(len(header),
		header...,
	))

	for i := 0; i <= 3; i++ {
		tmpJob := []fyne.CanvasObject{
			widget.NewLabel("jobId" + cast.ToString(i)),
			widget.NewLabel("开启"),
			widget.NewLabel("暂停"),
			container.NewHBox(widget.NewButton("Run", func() {
				fmt.Println("")
			}), widget.NewButton("Stop", func() {
				fmt.Println("")
			})),
		}
		jobList = append(jobList, container.NewGridWithColumns(len(tmpJob),
			tmpJob...,
		))
	}

	taskTab := container.NewBorder(
		nil,
		nil, nil, nil,
		container.NewVScroll(container.NewVBox(
			jobList...,
		)),
	)

	jobLists := []fyne.CanvasObject{
		widget.NewLabel("jobId"),
		widget.NewLabel("run"),
		widget.NewLabel("status"),
		widget.NewLabel("opt"),
	}
	jobLists2 := []fyne.CanvasObject{
		widget.NewLabel("jobId"),
		widget.NewLabel("run"),
		widget.NewLabel("status"),
		container.NewHBox(widget.NewButton("运行一次", func() {
			fmt.Println("")
		})),
	}

	scheduledTab := container.NewBorder(
		container.NewGridWithColumns(len(jobLists),
			jobLists...,
		),
		nil, nil, nil,
		container.NewVScroll(container.NewVBox(
			container.NewGridWithColumns(len(jobLists2),
				jobLists2...,
			)),
		),
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("常驻运行", taskTab),
		container.NewTabItem("定时运行", scheduledTab),
	)

	w.SetContent(tabs)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
