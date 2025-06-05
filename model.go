package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

type model struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
	state        string // "naming", "selecting", or "confirming"
	preview      string
	catalogName  string
	width        int
	height       int
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
		state:        "naming",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case m.state == "naming":
			switch msg.Type {
			case tea.KeyEnter:
				if m.catalogName != "" {
					m.list.Title = titleText + ": " + m.catalogName
					m.state = "selecting"
					return m, nil
				}
			case tea.KeyBackspace:
				if len(m.catalogName) > 0 {
					m.catalogName = m.catalogName[:len(m.catalogName)-1]
				}
			case tea.KeyRunes:
				m.catalogName += string(msg.Runes)
			}
			return m, nil

		case key.Matches(msg, m.keys.finalizeSelection):
			if m.state == "selecting" {
				catalog := generateOutputCatalog()
				data, err := yaml.Marshal(catalog)
				if err != nil {
					return m, tea.Println("Failed to generate preview: " + err.Error())
				}
				m.preview = string(data)
				m.state = "confirming"
				return m, nil
			}

		case m.state == "confirming":
			switch msg.String() {
			case "y", "Y":
				err := writeOutputCatalog("output.yaml")
				if err != nil {
					return m, tea.Println("Failed to write output.yaml: " + err.Error())
				}
				return m, tea.Quit
			case "n", "N":
				m.state = "selecting"
				return m, nil
			}
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// Minimum dimensions for showing border
	const minWidth = 80
	const minHeight = 24

	// Base content
	var content string
	if m.state == "naming" {
		content = lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.Styles.Title.Render(m.list.Title),
			"Enter catalog name: "+m.catalogName,
		)
	} else if m.state == "confirming" {
		content = lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.Styles.Title.Render(m.list.Title),
			"Preview of output catalog:",
			m.preview,
			"\nWrite to file? (Y/N)",
		)
	} else {
		content = m.list.View()
	}

	// Apply border if window is large enough
	if m.width >= minWidth && m.height >= minHeight {
		borderStyle := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#25A065"))
		content = borderStyle.Render(content)
	}

	return appStyle.Render(content)
}
