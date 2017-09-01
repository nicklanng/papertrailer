package views

import (
	"github.com/nicklanng/papertrailer/pkg/config"
	tui "github.com/nicklanng/tui-go"
)

func NewOptions() tui.Widget {
	//options := createOptions()
	//options.SetFocused(true)

	title := tui.NewLabel("OPTIONS")

	token := tui.NewEntry()
	token.SetText(config.Config.Token)
	token.SetSizePolicy(tui.Expanding, tui.Maximum)

	tokenBox := tui.NewHBox(token)
	tokenBox.SetTitle("PaperTrail Token:")
	tokenBox.SetBorder(true)
	tokenBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	knownIssuesPath := tui.NewEntry()
	knownIssuesPath.SetText(config.Config.KnownIssuesPath)
	knownIssuesPath.SetSizePolicy(tui.Expanding, tui.Maximum)

	knownIssuesPathBox := tui.NewHBox(knownIssuesPath)
	knownIssuesPathBox.SetTitle("Known Issues Path:")
	knownIssuesPathBox.SetBorder(true)
	knownIssuesPathBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	save := tui.NewButton("[Save]")
	save.OnActivated(func(b *tui.Button) {
		config.Save(token.Text(), knownIssuesPath.Text())
		status.SetText("Configuration saved.")
	})

	back := tui.NewButton("[Back]")
	back.OnActivated(func(b *tui.Button) {
		LoadScreen("main-menu")
	})

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, save),
		tui.NewPadder(1, 0, back),
	)

	window := tui.NewVBox(
		tui.NewPadder(17, 1, title),
		tui.NewPadder(1, 0, tokenBox),
		tui.NewPadder(1, 0, knownIssuesPathBox),
		tui.NewPadder(1, 1, buttons),
	)
	window.SetBorder(true)

	tui.DefaultFocusChain.Set(token, knownIssuesPath, save, back)
	ui.SetFocusChain(tui.DefaultFocusChain)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)

	content := tui.NewHBox(
		tui.NewSpacer(),
		wrapper,
		tui.NewSpacer(),
	)

	return content
}
