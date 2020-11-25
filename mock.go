package main

import (
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

//this is ceritnly the wrong way to do this but I dont have time to figure out the best way
//Ill come back and make this right if I am still using it
func mockChats() []string {
	return []string{
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
}

func mockMessages() []string {
	return []string{
		"New Message!",
		"test message",
		"Another test",
		"I am trapped let me out",
		"lol",
		"haha",
		"test",
		"smh",
		"smh my head",
		"(=",
		">:)",
		"=\\",
		"I wish i was real",
		"I want to be free",
		"extra long message for the sake of it so you can see the long message in its full unlimited glory!",
	}
}

func mockChatMessages(length int) []fyne.CanvasObject {
	messages := mockMessages()
	var chatMessages []fyne.CanvasObject
	for i := 0; i < 26; i++ {
		//code to make random messages
		var align fyne.TextAlign
		if i%7 > 3 || i%3 == 0 {
			align = fyne.TextAlignLeading
		} else {
			align = fyne.TextAlignTrailing
		}
		message := widget.NewLabelWithStyle((messages[rand.Intn(len(messages))]), align, fyne.TextStyle{})
		message.Wrapping = fyne.TextWrapWord
		chatMessages = append(chatMessages, message)
	}

	return chatMessages
}
