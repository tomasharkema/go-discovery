package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.treeview", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(int(code))
	}
}

func activate(app *gtk.Application) {
	win := gtk.NewApplicationWindow(app)
	win.SetTitle("Simple Treeview")
	win.SetDefaultSize(600, 300)

	// treeView := gtk.NewTreeView()

	// treeView.AppendColumn(createColumn("Name", NameColumn))
	// treeView.AppendColumn(createColumn("Comment", CommentColumn))

	// items := NewItemList()
	// treeView.SetModel(items)
	// win.SetChild(&treeView.Widget)

	// // Add some rows to the list store
	// items.Add("hello", "Gtk4")
	// items.Add("hello", "Gtk4")
	// items.Add("hello", "Gtk4")
	// items.Add("hello", "Gtk4")

	win.Show()
}
