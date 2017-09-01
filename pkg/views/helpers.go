package views

import tui "github.com/nicklanng/tui-go"

func createOptions(menuOptions []*menuOption) (options *tui.List) {
	options = tui.NewList()
	for _, o := range menuOptions {
		options.AddItems(o.text)
	}
	options.OnSelectionChanged(func(t *tui.List) {
		status.SetText(menuOptions[t.Selected()].description)
	})
	options.OnItemActivated(func(t *tui.List) {
		action := menuOptions[t.Selected()].action
		if action != nil {
			action()
		}
	})
	options.Select(0)

	return
}
