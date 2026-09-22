package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitApp() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.SetFixedSize(true)

	name := widget.NewLabel("");
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите ваше имя...")

	myWindow.SetContent(container.NewVBox(
		input,
		widget.NewButton("Отправить", func() {
			name.SetText(input.Text)
		}),
		name,
	))

	myWindow.ShowAndRun()
}