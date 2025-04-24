package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, model *list.Model) tea.Cmd {
		var title string

		if _, ok := model.SelectedItem().(item); !ok {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				if i, ok := model.SelectedItem().(item); ok {
					capabilityId := strings.Split(i.title, ":")[0]

					if _, ok := selectedCapabilities[capabilityId]; ok {
						if _, ok := triedToReselectCapability[capabilityId]; ok {
							return model.NewStatusMessage(statusMessageStyle("You can stop clicking on " + capabilityId))
						}
						triedToReselectCapability[capabilityId] = true
						return model.NewStatusMessage(statusMessageStyle("Already selected " + capabilityId))
					}
					selectedCapabilities[capabilityId] = i
					return model.NewStatusMessage(statusMessageStyle("Selected " + capabilityId))
				}

			case key.Matches(msg, keys.remove):
				index := model.Index()
				model.RemoveItem(index)
				if len(model.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}
				return model.NewStatusMessage(statusMessageStyle("Removed from canvas: " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose               key.Binding
	remove               key.Binding
	finishedCapabilities key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.remove,
		d.finishedCapabilities,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.remove,
			d.finishedCapabilities,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
		finishedCapabilities: key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "finish capabilities"),
		),
	}
}
