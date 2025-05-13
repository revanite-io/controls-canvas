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
	return &listKeyMap{
		finalizeSelection: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "generate output"),
		),
	}
}
