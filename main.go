package main

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
)

func main() {
	app := adw.NewApplication("io.harkema.go-discovery", 0)

	app.ConnectActivate(func() { activate(app) })

}

func activate(app *adw.Application) {
	win := adw.NewApplicationWindow(&app.Application)
	win.SetTitle("Simple Treeview")
	win.SetDefaultSize(600, 300)

	win.Present()
}
