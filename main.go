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

	app := gtk.NewApplication("com.example.test", 0)
	app.ConnectActivate(func() {
		t := newTest(ctx, app)
		t.Show()
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

type sTest struct {
	*gtk.Application
	window *gtk.ApplicationWindow
	header *gtk.HeaderBar
}

func newTest(ctx context.Context, app *gtk.Application) *sTest {
	t := sTest{Application: app}

	t.header = gtk.NewHeaderBar()

	t.window = gtk.NewApplicationWindow(app)
	t.window.SetTitle("Test")
	t.window.SetTitlebar(t.header)

	return &t
}

func (t *sTest) Show() {
	t.window.Show()
}
