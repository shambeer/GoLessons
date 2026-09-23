package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	myApp = app.New();
	myWindow = myApp.NewWindow("ShambeerApp");
)

func messageBox() {
			message := widget.NewLabel("67");
			message.Alignment = fyne.TextAlignCenter;
			dialog.ShowCustom("Внимание!", "Ok", message, myWindow);
		}

func InitApp() {
	myWindow.Resize(fyne.NewSize(400, 300));
	myWindow.SetFixedSize(true);

	myWindow.SetContent(container.NewVBox(
		widget.NewButton("Нажми меня", messageBox),
	));

	myWindow.ShowAndRun();
}