package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form")

	nameEntry := widget.NewEntry()
	//nameEntry.Resize(fyne.NewSize(10, 10))

	passEntry := widget.NewPasswordEntry()

	form := widget.NewForm(
		&widget.FormItem{Text: "Name", Widget: nameEntry},
		&widget.FormItem{Text: "Pass", Widget: passEntry},
	)
	form.OnSubmit = func() {
		fmt.Println("name:", nameEntry.Text, "pass:", passEntry.Text, "login in")
	}
	form.OnCancel = func() {
		fmt.Println("login canceled")
		myWindow.Close()
	}
	//form.Resize(fyne.NewSize(100, 100))

	myWindow.SetContent(form)
	//myWindow.Resize(fyne.Size{Width: 100, Height: 100})
	//myWindow.Resize(fyne.NewSize(50, 50))
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}
