package linux

import (
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/OlegLaban/sing_token/internal/adapter/config"
	"github.com/OlegLaban/sing_token/internal/adapter/logger"
	"github.com/OlegLaban/sing_token/internal/domain"
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
	w.SetContent(container.NewVBox(cont, input, widget.NewButton("Send", func() {
		cryptKey, err := generator.Generate(domain.NewPayload(input.Text, int(time.Now().Unix())))
		if err != nil {
			l.Error("can`t generate key", err)
			return
		}
		err = s.PutKey(cryptKey)
		if err != nil {
			l.Error("can`t share key\n", err)
		}
		cont.SetText(cryptKey)
		input.SetText("")
	})))
	w.ShowAndRun()
}
