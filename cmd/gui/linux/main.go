package main

import (
	"github.com/OlegLaban/sing_token/internal/pkg/app/gui/linux"
)

func main() {
	app := linux.NewAPP()
	app.Run("config/config.yml")
}
