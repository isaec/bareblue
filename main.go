package main

import (
	"fmt"
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

	messages := []string{
		"New Message, ",
		"test message... ",
		"Another test - ",
		"Message number ",
		"extra long message for the sake of it so you can see the long message in its full unlimited glory, number ",
	}

	chatHeader := widget.NewLabelWithStyle("Chat with noone", fyne.TextAlignCenter, fyne.TextStyle{true, false, true})

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
					widget.NewLabelWithStyle("Template Chat", fyne.TextAlignCenter, fyne.TextStyle{true, false, false}),
					widget.NewLabelWithStyle("last message (??/??)", fyne.TextAlignCenter, fyne.TextStyle{false, true, false})))
		},

		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Label).SetText(chats[id])
		})

	messageList.OnSelected = func(id widget.ListItemID) {
		chatHeader.SetText("Chat with " + chats[id])
	}
	//end messageList block

	messageEntry := widget.NewEntry()

	messageEntry.SetPlaceHolder("bareblue")

	var chatMessages []fyne.CanvasObject
	for i := 0; i < 26; i++ {
		//code to make random messages
		var align fyne.TextAlign
		if i%7 > 3 || i%3 == 0 {
			align = fyne.TextAlignLeading
		} else {
			align = fyne.TextAlignTrailing
		}
		chatMessages = append(chatMessages, widget.NewLabelWithStyle((messages[rand.Intn(len(messages))]+fmt.Sprint(i)), align, fyne.TextStyle{}))
	}

	chatContent := container.NewVScroll(widget.NewVBox(chatMessages...))

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(chatHeader, messageEntry, nil, nil), chatHeader, messageEntry, chatContent)

	messageAndContent := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, messageList, nil), messageList, content)

	w.SetContent(messageAndContent)
	w.Resize(fyne.NewSize(1200, 600))
	w.ShowAndRun()
}
