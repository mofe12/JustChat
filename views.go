package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

type chatView struct {
	historyScroll *tui.ScrollArea
}

type registerView struct{}

func views() {

	var currentView int

	chatView := chatView{}
	views := []tui.Widget{
		tui.NewVBox(tui.NewLabel("Press right arrow to continue ...")),
		chatView.chat(),
	}

	root := tui.NewVBox(views[0])
	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Left", func() {
		currentView = clamp(currentView-1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})

	ui.SetKeybinding("Right", func() {
		currentView = clamp(currentView+1, 0, len(views)-1)
		ui.SetWidget(views[currentView])
	})
	// SCROLLS UP FOR CHAT VIEW
	ui.SetKeybinding("Up", func() {
		chatView.historyScroll.SetAutoscrollToBottom(false)
		chatView.historyScroll.Scroll(0, -1)
	})
	// SCROLLS DOWN FOR CHAT VIEW
	ui.SetKeybinding("Down", func() { chatView.historyScroll.Scroll(0, 1) })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}