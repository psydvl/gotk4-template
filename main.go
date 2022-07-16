package main

import (
	"context"
	"os"
	"os/signal"

	_ "embed"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

//go:embed ui/main.ui
//TODO: try to use without global variable
var ui string

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
	var b *gtk.Builder
	b = gtk.NewBuilderFromString(ui, -1)
	this.Application = app

	this.header = gtk.NewHeaderBar()

	this.window = b.GetObject("main").Cast().(*gtk.ApplicationWindow)
	this.window.SetApplication(app)
}
