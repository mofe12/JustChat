package main

import (
	"github.com/marcusolsson/tui-go"
)

var logo = `
     ________  _____________________  ______________
    /__  _/ / / / ___/_  __/ ____/ /_/ / __  /_  __/
      / // / / /\__ \ / / / /   / __  / /_/ / / /
  ___/ // /_/ /___/ // / / /___/ / / / / / / / /
 /____//_____/_____//_/ /_____/_/ /_/_/ /_/ /_/
`

func (r *registerView) register() *tui.Box {
	user := tui.NewEntry()
	user.SetFocused(true)

	form := tui.NewGrid(0, 0)
	//form.SetBorder(true)
	form.AppendRow(tui.NewLabel("User:"), user)

	status := tui.NewStatusBar("Ready.")

	// login := tui.NewButton("[Login]")
	// login.OnActivated(func(b *tui.Button) {
	// 	status.SetText("Logged in.")
	// })

	// register := tui.NewButton("[Register]")

	// buttons := tui.NewHBox(
	// 	tui.NewSpacer(),
	// 	tui.NewPadder(1, 0, login),
	// 	tui.NewPadder(1, 0, register),
	// )

	window := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(14, 0, tui.NewLabel("Welcome to JustChat! Login or register.")),
		tui.NewPadder(1, 1, form),
		//buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(
		content,
		status,
	)

	tui.DefaultFocusChain.Set(user)

	return root
	// ui, err := tui.New(root)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ui.SetKeybinding("Esc", func() { ui.Quit() })

	// if err := ui.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
