package main

import (
	"fmt"
	"log"
	"time"

	"github.com/marcusolsson/tui-go"
)

type post struct {
	username string
	message  string
	time     time.Time
	userID   int
}

func (c *chatView) chat() *tui.Box {

	// CREATES SIDEBARS
	sidebar := tui.NewVBox(
		tui.NewLabel("CHANNELS"),
		tui.NewLabel("general"),
		tui.NewLabel("random"),
		tui.NewLabel(""),
		tui.NewLabel("DIRECT MESSAGES"),
		tui.NewLabel("slackbot"),
		tui.NewSpacer(),
	)
	sidebar.SetBorder(true)

	// CREATES THE HISTORY OF USERS TEXT
	history := tui.NewVBox()

	for _, m := range getNote() {
		history.Append(tui.NewHBox(
			tui.NewLabel(m.time.Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.username))),
			tui.NewLabel(m.message),
			tui.NewSpacer(),
		))
	}
	// SETS USERS TEXT AREA AND TEXT AT THE BOTTOM OF PAGE
	historyScroll := tui.NewScrollArea(history)
	c.historyScroll = historyScroll
	historyScroll.SetAutoscrollToBottom(true)

	historyScroll.ScrollToTop()
	//historyScroll.ScrollToTop()

	// CREATE VERTICAL BOX FOR USERS TEXT
	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	// CREATES, POSITIONS, AND GIVES SIZE POLICY OF INPUT
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	// CREATE HORIZONTAL BOX FOR USERS INPUT
	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	// CREATES A BOX TO CONTAIN BOTH USER HISTORY AND INPUT WITHOUT BORDERS
	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		historyScroll.SetAutoscrollToBottom(true)
		messages := &post{
			username: username,
			time:     time.Now(),
			message:  e.Text(),
			userID:   userid,
		}
		history.Append(tui.NewHBox(
			tui.NewLabel(messages.time.Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", messages.username))),
			tui.NewLabel(messages.message),
			tui.NewSpacer(),
		))
		input.SetText("")
		messages.storeMessage()
	})

	// PUTS EVEERYTHING INSIDE THE ROOT BOX
	root := tui.NewHBox(sidebar, chat)
	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Up", func() {
		historyScroll.SetAutoscrollToBottom(false)
		historyScroll.Scroll(0, -1)
	})
	ui.SetKeybinding("Down", func() { historyScroll.Scroll(0, 1) })
	return root
}
