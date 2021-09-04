package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

type chatView struct {
	historyScroll *tui.ScrollArea
}

type registerView struct {
	allowToNextPage bool
}

func views() {

	var currentView int

	chatView := chatView{}
	registerView := &registerView{}
	views := []tui.Widget{
		registerView.register(),
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
		registerView.allowToNextPage = false
	})

	ui.SetKeybinding("Right", func() {
		if registerView.allowToNextPage != false {
			currentView = clamp(currentView+1, 0, len(views)-1)
			ui.SetWidget(views[currentView])
		}
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
