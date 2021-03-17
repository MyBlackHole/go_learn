package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func setLog(ch chan string, mE *widget.Entry)  {
	var strAll string = ""
	for {
		str := <- ch
		strAll = str + "\n" + strAll
		mE.SetText(strAll)
	}
}

func setCh(ch chan string)  {
	var str string = ""
	for {
		ch <- str
		str += "ok"
		time.Sleep(6 * time.Second)
	}
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Entry")

	nameEntry := widget.NewEntry()
	//nameEntry.Resize(fyne.Size{Width: 200})
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		fmt.Println("name:", nameEntry.Text, "entered")
	}


	//passEntry := widget.NewPasswordEntry()
	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")
	//passEntry.Resize(fyne.NewSize(200, 200))

	//TextBox := container.NewVBox(widget.NewLabel("Name"), widget.NewLabel("Password"))
	//EntryBox := container.NewVBox(nameEntry, passEntry)
	//EntryBox.Resize(fyne.Size{Width: 200})

	//FromBox := container.NewHBox(TextBox, EntryBox)
	//print(FromBox)
	//FromBox.Resize(fyne.Size{Width: 200})

	//nameBox := container.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	//passwordBox := container.NewHBox(widget.NewLabel("Password"), layout.NewSpacer(), passEntry)
	ch := make(chan string)
	loginBtn := widget.NewButton("Login", nil)
	loginBtn.OnTapped = func() {
		go setCh(ch)
		//fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
		loginBtn.Disable()
	}

	exitBtn := widget.NewButton("exit", nil)
	exitBtn.OnTapped = func() {
		myWin.Close()
	}

	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.SetText("please enter\nyour description")
	multiEntry.SetText("please enter\nyour description")
	multiEntry.MultiLine = true

	go setLog(ch, multiEntry)

	content := container.NewVBox(nameEntry, passEntry, loginBtn, exitBtn, multiEntry)
	myWin.Resize(fyne.Size{Width: 300, Height: 300})
	myWin.SetFixedSize(true)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}
