package main

var ldb *LocalDatabase

var lorem = `Lorem ipsum dolor sit amet.`

func main() {
	ldb = &LocalDatabase{}
	ldb.init("./chat.db")

	chat()

	// box := tui.NewVBox(tui.NewLabel(lorem))
	// box.SetBorder(true)
	// box.SetSizePolicy(tui.Expanding, tui.Maximum)

	// root := tui.NewHBox()

	// //l := tui.NewLabel(lorem)

	// s := tui.NewScrollArea(box)

	// scrollBox := tui.NewVBox(s)
	// scrollBox.SetBorder(true)
	// scrollBox.SetSizePolicy(tui.Expanding, tui.Expanding)
	// root.Append(tui.NewVBox(tui.NewSpacer()))
	// root.Append(scrollBox)
	// root.Append(tui.NewVBox(tui.NewSpacer()))

	// ui, err := tui.New(root)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 0; i < 2; i++ {
	// 	fmt.Println("lol")
	// 	ui.SetKeybinding("Up", func() { s.Scroll(0, -1) })
	// }

	// ui.SetKeybinding("Esc", func() { ui.Quit() })
	// //ui.SetKeybinding("Up", func() { s.Scroll(0, -1) })
	// ui.SetKeybinding("Down", func() { s.Scroll(0, 1) })
	// ui.SetKeybinding("Left", func() { s.Scroll(-1, 0) })
	// ui.SetKeybinding("Right", func() { s.Scroll(1, 0) })
	// ui.SetKeybinding("a", func() { s.SetAutoscrollToBottom(true) })
	// ui.SetKeybinding("t", func() { s.ScrollToTop() })
	// ui.SetKeybinding("b", func() { s.ScrollToBottom() })

	// if err := ui.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
