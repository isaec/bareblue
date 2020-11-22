package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("bareblue")

	chats := []string{
		"Liam",
		"Olivia",
		"Noah",
		"Emma",
		"Oliver",
		"Ava",
		"William",
		"Sophia",
		"Elijah",
		"Isabella",
		"James",
		"Charlotte",
		"Benjamin",
		"Amelia",
		"Lucas",
		"Mia",
		"Mason",
		"Harper",
		"Ethan",
		"Evelyn"}

	messageList := widget.NewList(
		func() int { return len(chats) },

		func() fyne.CanvasObject {
			//groupChat := canvas.NewText(chats[chatIndex], color.White)
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.WarningIcon()), widget.NewLabel("Template Chat"))
		},

		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(chats[id])
		})

	messageList.OnSelected = func(id widget.ListItemID) {
		fmt.Println(id)
	}

	w.SetContent(messageList)
	w.Resize(fyne.NewSize(120, 240))
	w.ShowAndRun()
}
