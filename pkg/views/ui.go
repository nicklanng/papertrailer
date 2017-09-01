package views

import (
	tui "github.com/nicklanng/tui-go"
)

var (
	ui      tui.UI
	status  *tui.StatusBar
	screens map[string]func() tui.Widget
)

func init() {
	screens = map[string]func() tui.Widget{
		"main-menu":   NewMainMenu,
		"options":     NewOptions,
		"edit-groups": NewEditGroups,
	}
}

func InitializeUI() {
	status = tui.NewStatusBar("Ready.")

	root := tui.NewVBox(tui.NewHBox(), status)
	ui = tui.New(root)
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	selectedColor := tui.ColorBlue
	theme := tui.NewTheme()
	theme.SetStyle("list.item", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})
	theme.SetStyle("list.item.selected", tui.Style{Bg: selectedColor, Fg: tui.ColorWhite})
	theme.SetStyle("label.logo", tui.Style{Bg: tui.ColorDefault, Fg: selectedColor})
	theme.SetStyle("label.value", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})
	theme.SetStyle("button.focused", tui.Style{Bg: selectedColor, Fg: tui.ColorWhite})
	ui.SetTheme(theme)

	LoadScreen("main-menu")

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func LoadScreen(name string) {
	root := tui.NewVBox(screens[name](), status)
	ui.SetWidget(root)
}
