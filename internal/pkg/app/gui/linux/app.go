package linux

import (
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/OlegLaban/sing_token/internal/adapter/config"
	"github.com/OlegLaban/sing_token/internal/adapter/logger"
	keygenerator "github.com/OlegLaban/sing_token/internal/usecases/key_generator"
	"github.com/OlegLaban/sing_token/internal/usecases/share"
	"github.com/OlegLaban/sing_token/pkg/clipboard"
	"github.com/OlegLaban/sing_token/pkg/crypter"
)

type linuxAPP struct{}

func NewAPP() *linuxAPP {
	return &linuxAPP{}
}

func (la *linuxAPP) Run(configPath string) {
	a := app.New()
	w := a.NewWindow("Sign")
	input := widget.NewEntry()
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	conf, err := config.Parse(file)
	if err != nil {
		panic(err)
	}
	l := logger.New()
	cryp := crypter.New(conf.Crypto.Key)
	generator := keygenerator.NewGenerator(cryp, l)
	clipboard := clipboard.New()
	s := share.New(clipboard, l)
	cont := widget.NewLabel("some text")
	button := widget.NewButton("Send", func() {

	})

	w.SetContent(container.NewVBox(cont, input, button))
	w.ShowAndRun()
}
