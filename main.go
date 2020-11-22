package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
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

	//messageList block
	messageList := widget.NewList(
		func() int { return len(chats) },

		func() fyne.CanvasObject {

			contactIcon := canvas.NewImageFromFile("images/512px-Google_Contacts_icon.svg.png")
			contactIcon.FillMode = canvas.ImageFillContain
			contactIcon.SetMinSize(fyne.NewSize(40, 40))

			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
				contactIcon,
				fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
					widget.NewLabel("Template Chat"),
					widget.NewLabel("last message (??/??)")))
		},

		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(chats[id])
		})

	messageList.OnSelected = func(id widget.ListItemID) {
		fmt.Println(id)
	}
	//end messageList block

	messageEntry := widget.NewEntry()

	var chatMessages []fyne.CanvasObject
	for i := 0; i < 100; i++ {
		chatMessages = append(chatMessages, widget.NewLabel(string(i)))
	}

	chatContent := container.NewVScroll(widget.NewVBox(chatMessages...))

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, messageEntry, nil, nil), messageEntry, chatContent)

	messageAndContent := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, messageList, nil), messageList, content)

	w.SetContent(messageAndContent)
	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
