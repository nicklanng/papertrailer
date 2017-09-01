package views

import (
	tui "github.com/nicklanng/tui-go"
)

var logo = `   ___                      _             _ _
  / _ \__ _ _ __   ___ _ __| |_ _ __ __ _(_) | ___ _ __
 / /_)/ _' | ,_ \ / _ \ '__| __| '__/ _, | | |/ _ \ '__|
/ ___/ (_| | |_) |  __/ |  | |_| | | (_| | | |  __/ |
\/    \__,_| .__/ \___|_|   \__|_|  \__,_|_|_|\___|_|
           |_|                                          `

func NewMainMenu() tui.Widget {
	menuOptions := []*menuOption{
		&menuOption{
			text:        "Watch Groups (Not implemented)",
			description: "View Papertrail groups",
		},
		&menuOption{
			text:        "Edit Groups (WIP)",
			description: "Edit Papertrail groups",
			action: func() {
				LoadScreen("edit-groups")
			},
		},
		&menuOption{
			text: "",
		},
		&menuOption{
			text:        "Options",
			description: "Set up your Papertrail token and known issues list",
			action: func() {
				LoadScreen("options")
			},
		},
		&menuOption{
			text: "",
		},
		&menuOption{
			text:        "Exit",
			description: "Close Papertrailer",
			action: func() {
				ui.Quit()
			},
		},
	}

	options := createOptions(menuOptions)
	options.SetFocused(true)

	logo := tui.NewLabel(logo)
	logo.SetStyleName("logo")

	window := tui.NewVBox(
		tui.NewPadder(1, 1, logo),
		tui.NewPadder(17, 0, tui.NewLabel("Welcome to Papertrailer!")),
		tui.NewPadder(1, 1, options),
	)
	window.SetBorder(true)

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
