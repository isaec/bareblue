package main

import (
	"math/rand"
	"time"

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

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	chats := mockChats()

	chatHeader := widget.NewLabelWithStyle("select a chat", fyne.TextAlignCenter, fyne.TextStyle{true, false, true})

	//messageList block
	messageList := widget.NewList(
		func() int { return len(chats) },

		func() fyne.CanvasObject {

			contactIcon := canvas.NewImageFromFile("images/512px-Google_Contacts_icon.svg.png")
			contactIcon.FillMode = canvas.ImageFillContain
			contactIcon.SetMinSize(fyne.NewSize(40, 40))

			lastMessage := widget.NewLabelWithStyle("last message lalalalalaalalalal", fyne.TextAlignCenter, fyne.TextStyle{false, true, false})
			lastMessage.Wrapping = fyne.TextTruncate

			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
				contactIcon,
				fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
					widget.NewLabelWithStyle("Template Chat", fyne.TextAlignCenter, fyne.TextStyle{true, false, false}),
					lastMessage))
		},

		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(chats[id])
		})

	messageList.OnSelected = func(id widget.ListItemID) {
		chatHeader.SetText(chats[id])
	}
	//end messageList block

	messageEntry := widget.NewEntry()

	messageEntry.SetPlaceHolder("bareblue")

	chatMessages := mockChatMessages(25)

	chatContent := container.NewVScroll(widget.NewVBox(chatMessages...))

	chatContent.ScrollToBottom()

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(chatHeader, messageEntry, nil, nil), chatHeader, messageEntry, chatContent)

	messageAndContent := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, messageList, nil), messageList, content)

	w.SetContent(messageAndContent)
	w.Resize(fyne.NewSize(1200, 600))
	addMenus(&a, &w)
	w.ShowAndRun()
}
