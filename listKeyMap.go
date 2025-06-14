package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type listKeyMap struct {
	list.KeyMap
	finalizeSelection key.Binding
}

func newListKeyMap() *listKeyMap {
	km := &listKeyMap{
		KeyMap: list.KeyMap{
			CursorUp: key.NewBinding(
				key.WithKeys("up"),
				key.WithHelp("↑", "up"),
			),
			CursorDown: key.NewBinding(
				key.WithKeys("down"),
				key.WithHelp("↓", "down"),
			),
			Filter: key.NewBinding(
				key.WithKeys("/"),
				key.WithHelp("/", "filter"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
			// Disable other key bindings
			GoToStart:   key.NewBinding(),
			GoToEnd:     key.NewBinding(),
			NextPage:    key.NewBinding(),
			PrevPage:    key.NewBinding(),
			ClearFilter: key.NewBinding(),
		},
		finalizeSelection: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "generate output"),
		),
	}

	return km
}

// ShortHelp returns the help text we want to show
func (k listKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.CursorUp,
		k.CursorDown,
		k.Filter,
		k.Quit,
		k.finalizeSelection,
	}
}

// FullHelp returns the help text we want to show
func (k listKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.CursorUp,
			k.CursorDown,
			k.Filter,
			k.Quit,
			k.finalizeSelection,
		},
	}
}
