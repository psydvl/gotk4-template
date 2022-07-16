package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var instance sApplication

	app := gtk.NewApplication("com.example.test", 0)
	app.ConnectActivate(func() {
		instance.init(ctx, app)
		instance.Show()
	})

	go func() {
		<-ctx.Done()
		glib.IdleAdd(app.Quit)
	}()

	if code := app.Run(os.Args); code > 0 {
		cancel()
		os.Exit(code)
	}
}

type sApplication struct {
	*gtk.Application
	window *gtk.ApplicationWindow
	header *gtk.HeaderBar
}

func (this *sApplication) Show() {
	this.window.Show()
}

func (this *sApplication) init(ctx context.Context, app *gtk.Application) {
	this.Application = app

	this.header = gtk.NewHeaderBar()

	layout := gtk.NewBox(gtk.OrientationVertical, 0)
	layout.SetMarginStart(10)
	layout.SetMarginTop(10)
	layout.SetMarginEnd(10)
	layout.SetMarginBottom(10)
	layout.Append(gtk.NewLabel("Hello, world!"))

	this.window = gtk.NewApplicationWindow(app)
	this.window.SetTitle("Test")
	this.window.SetTitlebar(this.header)
	this.window.SetChild(layout)
}
