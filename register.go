package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

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
		userText := e.Text()
		boolean, userid := checkForID(userText)
		log.Println(userid)
		currentUser := map[int]string{}
		/* gets user when id is in the string
		example: mofe#123 */
		if boolean == true {
			currentUser = getCurrentUser(userid, boolean)
		}
		if boolean == false {
			createUser(e.Text())
			user.SetText("")
			currentUser = getCurrentUser(userid, boolean)
		}
		for k, v := range currentUser {
			globUsername = v
			globUserid = k
		}
		for _, m := range getNote() {
			history.Append(tui.NewHBox(
				tui.NewLabel(m.time.Format("15:04")),
				tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", m.globUsername))),
				tui.NewLabel(m.message),
				tui.NewSpacer(),
			))
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

func checkForID(userText string) (bool, int) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	if re.MatchString(userText) != false {
		submatchall := re.FindAllString(userText, -1)
		for _, element := range submatchall {
			id, err := strconv.Atoi(element)
			if err != nil {
				log.Printf("Could not get id from string %v\n", err)
			}
			return re.MatchString(userText), id
		}
	}
	return re.MatchString(userText), -1

}
