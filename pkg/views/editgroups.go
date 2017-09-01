package views

import (
	tui "github.com/nicklanng/tui-go"
)

func NewEditGroups() tui.Widget {
	menuOptions := []*menuOption{
		&menuOption{
			text:        "Create a new group",
			description: "Create a new group",
			action: func() {

			},
		},
		&menuOption{
			text: "",
		},
		&menuOption{
			text:        "Back to Main Menu",
			description: "Back to Main Menu",
			action: func() {
				LoadScreen("main-menu")
			},
		},
	}

	options := createOptions(menuOptions)
	options.SetFocused(true)

	window := tui.NewVBox(
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
