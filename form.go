package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type formField struct {
	key         string
	value       string
	description string
	isEditing   bool
}

func (i formField) Title() string {
	if i.isEditing {
		return i.key + ": " + i.value + "_"
	}
	return i.key + ": " + i.value
}

func (i formField) Description() string { return i.description }
func (i formField) FilterValue() string { return i.key }

func newFormFields() []list.Item {
	return []list.Item{
		formField{
			key:         "id",
			description: "A unique identifier for this catalog",
		},
		formField{
			key:         "title",
			description: "The display name for this catalog",
		},
		formField{
			key:         "description",
			description: "A detailed description of this catalog's purpose",
		},
		formField{
			key:         "version",
			description: "The version number of this catalog",
		},
		formField{
			key:         "last-modified",
			description: "The date this catalog was last modified",
		},
	}
}

func getFormStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
} 