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
	state        string
	preview      string
	width        int
	height       int
	selectedUrls []string
	descWidth    int
	sizeWarning  string
}

type catalogItem struct {
	title       string
	description string
	urls        []string
}

func (i catalogItem) Title() string       { return i.title }
func (i catalogItem) Description() string { return i.description }
func (i catalogItem) FilterValue() string { return i.title }

func newCatalogInputModel() model {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	items := []list.Item{
		catalogItem{
			title:       "Common Cloud Controls",
			description: "Default catalog with cloud security controls",
			urls: []string{
				"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/controls.yaml",
				"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/threats.yaml",
				"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/capabilities.yaml",
			},
		},
		catalogItem{
			title:       "Future reference options will be added here",
			description: "(Selecting this placeholder will just close the program)",
			urls:        []string{},
		},
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	catalogCanvas := list.New(items, delegate, 0, 0)
	catalogCanvas.Title = "Select Catalog"
	catalogCanvas.Styles.Title = titleStyle

	return model{
		list:         catalogCanvas,
		keys:         listKeys,
		delegateKeys: delegateKeys,
		state:        "catalog",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	currentModel = m

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		
		m.descWidth = (msg.Width-h)/2 - 10
		
		const minWidth = 60
		if msg.Width < minWidth {
			m.sizeWarning = "Window too small. Please resize to view content."
			m.descWidth = 0
		} else {
			m.sizeWarning = ""
			if m.state == "selecting" {
				choices := loadChoicesWithUrls(m.selectedUrls)
				m.list.SetItems(choices)
			}
		}

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case m.state == "catalog":
			switch msg.Type {
			case tea.KeyEnter:
				if item, ok := m.list.SelectedItem().(catalogItem); ok {
					m.selectedUrls = item.urls
					if item.title == "Common Cloud Controls" {
						choices := loadChoicesWithUrls(item.urls)
						m.list.SetItems(choices)
						m.list.Title = titleText
						m.state = "naming"
					} else {
						return m, tea.Quit
					}
				}
				return m, nil
			case tea.KeyUp, tea.KeyDown:
				newListModel, cmd := m.list.Update(msg)
				m.list = newListModel
				cmds = append(cmds, cmd)
				return m, tea.Batch(cmds...)
			}
			return m, nil

		case m.state == "naming":
			switch msg.Type {
			case tea.KeyEnter:
				if catalogName != "" {
					m.list.Title = titleText + ": " + catalogName
					m.state = "selecting"
					return m, nil
				}
			case tea.KeyBackspace:
				if len(catalogName) > 0 {
					catalogName = catalogName[:len(catalogName)-1]
				}
			case tea.KeyRunes:
				catalogName += string(msg.Runes)
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

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	const minWidth = 80
	const minHeight = 24
	const twoColumnWidth = 120

	if m.sizeWarning != "" {
		return m.sizeWarning
	}

	var content string
	if m.state == "catalog" {
		content = m.list.View()
	} else if m.state == "naming" {
		content = lipgloss.JoinVertical(
			lipgloss.Left,
			m.list.Styles.Title.Render(m.list.Title),
			"Enter catalog name: "+catalogName,
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
		if m.width >= twoColumnWidth {
			catalog := generateOutputCatalog()
			data, err := yaml.Marshal(catalog)
			if err != nil {
				data = []byte("Error generating catalog preview")
			}

			listWidth := (m.width * 3) / 5
			previewWidth := (m.width * 2) / 5 - 4

			listContainer := lipgloss.NewStyle().
				Width(listWidth).
				Height(m.height - 4).
				Render(m.list.View())

			catalogPreview := lipgloss.NewStyle().
				Width(previewWidth).
				Height(m.height - 4).
				Render(string(data))

			content = lipgloss.JoinHorizontal(
				lipgloss.Top,
				listContainer,
				lipgloss.NewStyle().PaddingLeft(2).Render(catalogPreview),
			)
		} else {
			content = m.list.View()
		}
	}

	contentStyle := lipgloss.NewStyle().
		Width(m.width - 4).
		Height(m.height - 4)

	if m.width >= minWidth && m.height >= minHeight {
		content = appStyle.Render(contentStyle.Render(content))
	} else {
		content = contentStyle.Render(content)
	}

	return content
}

var currentModel interface{}
