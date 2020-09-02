package main

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("MultineLine")

	label := widget.NewLabel("Birşeyler yaz , butona bas bildirim gelecek :)")
	lines := widget.NewMultiLineEntry()

	bildirim := func() {
		if strings.Trim(lines.Text, " ") == "" {
			label.SetText("Boş elemanı nasıl yazayım değil mi :))")
			return
		}
		label.SetText(lines.Text)
		uyg.SendNotification(fyne.NewNotification("Butona Bastın", lines.Text))

	}
	buton := widget.NewButton("Bas kardeşim", bildirim)

	group := widget.NewGroup("Bildirim al ", label, lines, buton)

	win.SetContent(group)

	win.ShowAndRun()
}
