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

func (t *sTest) Show() {
	t.window.Show()
}

func newTest(ctx context.Context, app *gtk.Application) *sTest {
	var b *gtk.Builder
	b = gtk.NewBuilderFromString(ui, -1)
	t := sTest{Application: app}

	t.header = gtk.NewHeaderBar()

	t.window = b.GetObject("main").Cast().(*gtk.ApplicationWindow)
	t.window.SetApplication(app)

	return &t
}
