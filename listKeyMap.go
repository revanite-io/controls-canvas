package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type listKeyMap struct {
	list.KeyMap
	finalizeSelection key.Binding
	makeSelection     key.Binding
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
			Quit: key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
			// Disable unused bindings
			GoToStart:   key.NewBinding(),
			GoToEnd:     key.NewBinding(),
			NextPage:    key.NewBinding(),
			PrevPage:    key.NewBinding(),
			Filter:      key.NewBinding(),
			ClearFilter: key.NewBinding(),
		},
		makeSelection: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		finalizeSelection: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "continue"),
		),
	}

	return km
}

// ShortHelp returns the help text we want to show
func (k listKeyMap) ShortHelp() []key.Binding {
	if currentModel != nil {
		if m, ok := currentModel.(model); ok {
			switch m.state {
			case "selecting":
				return []key.Binding{
					k.makeSelection,
					k.finalizeSelection,
					key.NewBinding(
						key.WithKeys("backspace"),
						key.WithHelp("backspace", "deselect"),
					),
				}
			case "naming":
				return []key.Binding{
					k.makeSelection,
				}
			default: // catalog state
				return []key.Binding{
					k.makeSelection,
				}
			}
		}
	}
	return []key.Binding{
		k.makeSelection,
	}
}

// FullHelp returns the help text we want to show
func (k listKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
