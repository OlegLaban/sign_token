package clipboard

import "golang.design/x/clipboard"

type clip struct{}

func New() *clip {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	return &clip{}
}

func (c *clip) Send(data []byte) {
	clipboard.Write(clipboard.FmtText, data)
}
