package main

import "github.com/charmbracelet/bubbles/key"

type listKeyMap struct {
	toggleSpinner     key.Binding
	toggleHelpMenu    key.Binding
	finalizeSelection key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
		finalizeSelection: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "generate output"),
		),
	}
}
