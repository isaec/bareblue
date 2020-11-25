package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/cmd/fyne_settings/settings"
	"fyne.io/fyne/dialog"
)

func addMenus(a *fyne.App, w *fyne.Window) {
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := (*a).NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480*1.5))
		w.CenterOnScreen()
		w.Show()
	})

	helpItem := fyne.NewMenuItem("Get help", func() {
		dialog.ShowConfirm("alert", "there is no help", nil, *w)
	})

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Options", settingsItem),
		fyne.NewMenu("Help", helpItem),
	)
	(*w).SetMainMenu(mainMenu)
}
