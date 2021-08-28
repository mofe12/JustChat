package main

import (
	"github.com/marcusolsson/tui-go"
)

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

	user.OnSubmit(func(e *tui.Entry) {
		createUser(e.Text())
		user.SetText("")
		user := getCurrentUser()
		for k, v := range user {
			username = v
			userid = k
		}
	})
	window := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(14, 0, tui.NewLabel("Welcome to JustChat! Login or register.")),
		tui.NewPadder(1, 1, form),
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
}
