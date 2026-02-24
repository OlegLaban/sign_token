package linux

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type linuxAPP struct{}

func NewAPP() *linuxAPP {
	return &linuxAPP{}
}

func (la *linuxAPP) Run() {
	a := app.New()
	w := a.NewWindow("Sign")
	input := widget.NewEntry()

	w.SetContent(container.NewVBox(widget.NewLabel("some text"), input, widget.NewButton("Send", func() {
		log.Println("Content was:", input.Text)
		input.SetText("")
	})))
	w.ShowAndRun()
}
