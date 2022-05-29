package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setLog(ch chan string, mE *widget.Entry) {
	var strAll string = ""
	for {
		str := <-ch
		strAll = str + "\n" + strAll
		mE.SetText(strAll)
	}
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("VPN")

	// 用户名
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")

	// 密码
	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	ch := make(chan string)

	loginBtn := widget.NewButton("Login", nil)
	loginBtn.OnTapped = func() {
		go mainLoginVPN(ch, nameEntry.Text, passEntry.Text)
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
		loginBtn.Disable()
	}

	exitBtn := widget.NewButton("exit", nil)
	exitBtn.OnTapped = func() {
		myWin.Close()
	}

	multiEntry := widget.NewEntry()
	multiEntry.MultiLine = true
	multiEntry.Disable()

	go setLog(ch, multiEntry)

	content := container.NewVBox(nameEntry, passEntry, loginBtn, exitBtn, multiEntry)
	myWin.Resize(fyne.Size{Width: 300, Height: 300})
	myWin.SetFixedSize(true)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}

// go build -ldflags="-r ./lib/" login.go ss.go ukey.go config.go