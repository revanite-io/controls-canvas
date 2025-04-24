package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type catalogInputModel struct {
	textInput textinput.Model
}

func newNameInputModel() catalogInputModel {
	ti := textinput.New()
	ti.Placeholder = "Enter catalog name"
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 30

	return catalogInputModel{textInput: ti}
}

func (m catalogInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m catalogInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			catalogName = m.textInput.Value()
			return newCatalogInputModel(), nil // 🔁 transition to the selection model
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m catalogInputModel) View() string {
	return appStyle.Render("Name your catalog:\n\n" + m.textInput.View() + "\n\nPress Enter to continue")
}
