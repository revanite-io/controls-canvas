package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func newCatalogInputModel() model {
	var (
		// itemGenerator randomItemGenerator
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	items := loadChoices()

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	catalogCanvas := list.New(items, delegate, 0, 0)
	catalogCanvas.Title = titleText
	catalogCanvas.Styles.Title = titleStyle

	return model{
		list:         catalogCanvas,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.finalizeSelection):
			catalog = generateOutputCatalog()

			// TODO: display catalog and prompt user for output file name

			err := writeOutputCatalog(catalog, "output.yaml")
			if err != nil {
				return m, tea.Println("Failed to write output.yaml: " + err.Error())
			}
			return m, tea.Quit
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}
